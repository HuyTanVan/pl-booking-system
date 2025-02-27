APP_NAME = python
private_proto:
	python -m grpc_tools.protoc -I ./protos \
	  --python_out=./pb \
	  --grpc_python_out=./pb \
	  ./protos/rpc_update_ticket_prices.proto
