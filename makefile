# #* GET FILE ENV
# include .env
# export $(shell sed 's/=.*//' .env)
APP_NAME = server
run:
	go run ./cmd/${APP_NAME}/

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:123b123b@localhost:5433/booking_ticket_db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:123b123b@localhost:5433/booking_ticket_db?sslmode=disable" -verbose down

proto_path:
	export PATH="$PATH:$(go env GOPATH)/bin"
		
public_proto:
	rm -f internal/pb/public_proto/*.go
	protoc -I ./internal/proto \
	--go_out ./internal/pb --go_opt=paths=source_relative \
    --go-grpc_out ./internal/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ./internal/pb --grpc-gateway_opt=paths=source_relative \
   	internal/proto/public_proto/*.proto

private_proto:
	rm -f internal/pb/private_proto/*.go
	protoc -I ./internal/proto \
	--go_out ./internal/pb --go_opt=paths=source_relative \
    --go-grpc_out ./internal/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ./internal/pb --grpc-gateway_opt=paths=source_relative \
   	internal/proto/private_proto/*.proto