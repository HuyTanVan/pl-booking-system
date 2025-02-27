package public_grpc

import (
	"context"
	"fmt"
	"net/mail"

	db "plbooking_go_structure1/internal/db/sqlc"
	pb "plbooking_go_structure1/internal/pb/public_proto"

	worker "plbooking_go_structure1/internal/redis_workers"
	"plbooking_go_structure1/internal/utils"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// request and reponse are defined in pb package as grpc format
func (server *PublicGrpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println(req.GetEmail())
	// Check if email is valid
	email, err := mail.ParseAddress(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid email address%s", err)
	}
	// compares passwords
	if req.GetPassword() != req.GetRepeatPassword() {
		return nil, status.Errorf(codes.InvalidArgument, "password and repeat password does not match")
	}
	hashedPass, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password%s", err)
	}

	// create token for email verification
	token, err := utils.RandomString(32)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create verification token%s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Email:    email.Address,
			Password: hashedPass,
		},
		CreateEmailVerificationTokenParams: db.CreateEmailVerificationTokenParams{
			Token: token,
		},
		AfterCreate: func(user db.User, token db.EmailVerificationToken) error {
			taskPay := &worker.PayloadSendVerifyEmail{
				Email: user.Email,
			}
			// give the task to distributor to process after done insert user into db
			return server.TaskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPay)
		},
	}

	user, err := server.Pgdbc.CreateUserTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "email already registered: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to register account: %s", err)
	}

	rsp := &pb.CreateUserResponse{User: &pb.User{Email: user.User.Email, CreatedAt: timestamppb.Now()}}
	return rsp, nil
}
