package initializer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"plbooking_go_structure1/global"
	"plbooking_go_structure1/internal/grpc/private_grpc"
	"plbooking_go_structure1/internal/grpc/public_grpc"
	pb "plbooking_go_structure1/internal/pb/public_proto"
	worker "plbooking_go_structure1/internal/redis_workers"
	"plbooking_go_structure1/internal/rest"
	"plbooking_go_structure1/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
)

// HTTP(REST) Server
func InitHttpServer(distributer worker.TaskDistributor) (*rest.HttpServer, error) {
	var r *gin.Engine
	if global.Config.HTTPServer.Mode == "development" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else if global.Config.HTTPServer.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		panic("failed to initialize routers-> modes allowed: 'development'/'production'")
	}

	tMaker, err := token.NewJWTMaker(global.Config.JWTToken.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create jwt token maker: %w", err)
	}
	return &rest.HttpServer{
		Config:          global.Config,
		Pgdbc:           global.Pgdbc,
		Rdb:             global.Rdb,
		Router:          r,
		TaskDistributor: distributer,
		Token:           tMaker}, nil
}

// PUBLIC GRPC SERVER
func InitPrivateGrpcServer(distributer worker.TaskDistributor) (*private_grpc.PrivateGrpcServer, error) {
	tokenMaker, err := token.NewJWTMaker(global.Config.JWTToken.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &private_grpc.PrivateGrpcServer{
		Config:          global.Config,
		Pgdbc:           global.Pgdbc,
		Rdb:             global.Rdb,
		Token:           tokenMaker,
		TaskDistributor: distributer,
	}, nil
}

// PUBLIC GRPC SERVER
func InitPublicGrpcServer(distributer worker.TaskDistributor) (*public_grpc.PublicGrpcServer, error) {
	tokenMaker, err := token.NewJWTMaker(global.Config.JWTToken.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &public_grpc.PublicGrpcServer{
		Config:          global.Config,
		Pgdbc:           global.Pgdbc,
		Rdb:             global.Rdb,
		Token:           tokenMaker,
		TaskDistributor: distributer,
	}, nil
}

func StartPublicGrpcGateway(distributer worker.TaskDistributor, ctx context.Context, waitGroup *errgroup.Group) error {

	server, _ := InitPublicGrpcServer(distributer)
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	err := pb.RegisterPremierLeagueBookingHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	// Set up CORS using the "rs/cors" package
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Your frontend origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With"},
		Debug:          true,
	})
	handler := corsHandler.Handler(mux)
	// declare http server to gracefully shutdown later
	httpServer := &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%d", server.Config.GRPCGatewayServer.Port),
	}
	waitGroup.Go(func() error {
		log.Printf("start HTTP gateway to GRPC at %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Println("HTTP gateway to GRPC failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("graceful shutdown HTTP gateway to GRPC server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Println("failed to shutdown HTTP gateway server")
			return err
		}

		log.Println("HTTP gateway server is stopped")
		return nil
	})
	return nil
}
