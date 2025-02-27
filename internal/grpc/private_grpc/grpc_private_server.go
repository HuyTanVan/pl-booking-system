package private_grpc

import (
	"context"
	"log"
	"net"
	worker "plbooking_go_structure1/internal/redis_workers"
	"plbooking_go_structure1/pkg/setting"

	db "plbooking_go_structure1/internal/db/sqlc"
	pb "plbooking_go_structure1/internal/pb/private_proto"

	"plbooking_go_structure1/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// PublicServer serves HTTP requests from clients
type PrivateGrpcServer struct {
	pb.UnimplementedPrivatePremierLeagueBookingServer
	Config          setting.Config
	Pgdbc           db.Store
	Rdb             *redis.Client
	Router          *gin.Engine
	Token           token.IMaker
	TaskDistributor worker.TaskDistributor
}

func (server *PrivateGrpcServer) StartServer(ctx context.Context, waitGroup *errgroup.Group) {
	grpcServer := grpc.NewServer()
	pb.RegisterPrivatePremierLeagueBookingServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":9159")

	if err != nil {
		log.Fatal("cannot create listener of grpc:", err)
	}
	waitGroup.Go(func() error {
		log.Printf("start Private GPRC server at %s", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("cannot serve the listener to public GRPC server:", err)
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("graceful shutdown gRPC server")

		grpcServer.GracefulStop()
		log.Println("gRPC server is stopped")

		return nil
	})
}
