package public_grpc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "plbooking_go_structure1/internal/db/sqlc"
	pb "plbooking_go_structure1/internal/pb/public_proto"
	worker "plbooking_go_structure1/internal/redis_workers"
	"plbooking_go_structure1/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *PublicGrpcServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.Pgdbc.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "invalid email or password:%s", err)
		}
		return nil, status.Errorf(codes.Internal, "something went wrong!")
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "invalid email or password:%s", err)
	}
	// if email and password correct, then check if account is activated yet.
	if !user.IsActive {
		err := server.TaskDistributor.DistributeTaskSendVerifyEmail(ctx, &worker.PayloadSendVerifyEmail{
			Email: user.Email})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "account is not activated and cannot send email verification:%s", err)
		}
		return nil, status.Errorf(codes.OK, "email verification has been sent")
	}
	// create access token
	dur, err := time.ParseDuration(server.Config.JWTToken.AccessTokenDuration)
	if err != nil {
		err := fmt.Errorf("can not parse time")
		return nil, err
	}
	accessTok, accessPay, err := server.Token.CreateToken(user.Email, dur)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create access token")
	}
	// create refresh token
	refreshTok, refreshPay, err := server.Token.CreateToken(user.Email, dur)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create refresh token")
	}
	// create session
	session, err := server.Pgdbc.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPay.ID,
		Email:        user.Email,
		RefreshToken: refreshTok,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPay.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create session")
	}
	rsp := &pb.LoginUserResponse{
		User:                  &pb.User{Email: user.Email, CreatedAt: timestamppb.New(user.CreatedAt)},
		SessionId:             session.ID.String(),
		AccessToken:           accessTok,
		RefreshToken:          refreshTok,
		AccessTokenExpiresAt:  timestamppb.New(accessPay.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPay.ExpiredAt),
	}
	return rsp, nil
}
