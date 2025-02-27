package public_grpc

import (
	"context"
	"fmt"
	"strings"

	"plbooking_go_structure1/internal/token"

	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *PublicGrpcServer) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	values := md.Get(authorizationHeader)
	if len(values) == 0 {

	}
	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}
	// check auth type = Bearer
	if authType := strings.ToLower(fields[0]); authType != authorizationHeader {
		return nil, fmt.Errorf("unspported authorization header: %s", authType)
	}
	accessTok := fields[1]
	payload, err := server.Token.VerifyToken(accessTok)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}
