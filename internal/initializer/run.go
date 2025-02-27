package initializer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func Run() {

	LoadConfig()
	InitRedis()       // global Rdb is assigned
	InitPostgreSQLC() // global Pgdbc is assigned
	// InitKafka()
	distributer := InitRedisTaskDistributer()
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// init HTTP server
	httpS, err := InitHttpServer(*distributer)
	if err != nil {
		log.Fatal("failed to initialize HTTP Server->", err)

	}
	grpcP, err := InitPrivateGrpcServer(*distributer)
	if err != nil {
		log.Fatal("failed to initialize Private GRPC Server->", err)
	}
	// init public GRPC server
	grpcS, err := InitPublicGrpcServer(*distributer)
	if err != nil {
		log.Fatal("failed to initialize Grpc Server->", err)

	}

	// grpcG, err := InitGrpcGatewayServer()
	// if err != nil {
	// 	log.Fatal("failed to initialize Grpc gateway Server->", err)

	// }
	waitGroup, ctx := errgroup.WithContext(ctx)
	// init redis background task processor
	taskPro := InitRedisTaskProccessor()
	go taskPro.Start()
	grpcP.StartServer(ctx, waitGroup)
	grpcS.StartServer(ctx, waitGroup)
	StartPublicGrpcGateway(*distributer, ctx, waitGroup)
	go httpS.StartServer()

	err = waitGroup.Wait()
	if err != nil {
		log.Fatalf("error from wait group->>%v", err)
	}

}

//  kafka initializer code

// func Run() {
// 	InitKafka()
// }
