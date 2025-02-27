package private_grpc

import (
	"context"
	pb "plbooking_go_structure1/internal/pb/private_proto"
)

// update ticket price based on the ticket availability
func (server *PrivateGrpcServer) UpdateTicketPrices(ctx context.Context, req *pb.UpdateTicketPricesRequest) (*pb.UpdateTicketPricesResponse, error) {
	// 1. get original price.

	// 2. get ticket available.
	// 3. get match date
	// 4. get seat
	// 5. get teams

	return &pb.UpdateTicketPricesResponse{Status: true, Message: "price updates have been sent to Kafka Server"}, nil
}
