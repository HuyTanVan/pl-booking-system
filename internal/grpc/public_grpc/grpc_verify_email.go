package public_grpc

import (
	"context"
	"database/sql"

	db "plbooking_go_structure1/internal/db/sqlc"
	pb "plbooking_go_structure1/internal/pb/public_proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// this function handles the verify email request from user since they had a activation account link
func (server *PublicGrpcServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	user, err := server.Pgdbc.GetUser(ctx, req.GetUserId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not in record:%s", err)
		}
		return nil, status.Errorf(codes.Internal, "something went wrong!")
	}

	if !user.IsActive {
		// check token, if user owns token, then allows to update user status
		token, err := server.Pgdbc.GetEVTokenByUserID(ctx, user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Errorf(codes.NotFound, "token doesnt belong to user:%s", err)
			}
			return nil, status.Errorf(codes.Internal, "something went wrong!")
		}
		if token.Token != req.GetToken() {
			return nil, status.Errorf(codes.Internal, "token verification doesnt match!")
		}

		a := server.Pgdbc.UpdateUser(ctx, db.UpdateUserParams{ID: user.ID, IsActive: sql.NullBool{Bool: true, Valid: true}})
		if a != nil {
			return nil, status.Errorf(codes.Internal, "cannot update user status:%s", err)
		}
	} else {
		return nil, status.Errorf(codes.Internal, "this email address was activated")
	}

	// if user is not activated in db, then update it to true

	rsp := &pb.VerifyEmailResponse{IsVerified: true}
	return rsp, nil
}
