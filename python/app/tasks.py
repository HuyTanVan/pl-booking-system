from pb import rpc_update_ticket_prices_pb2 as pb2
from pb import rpc_update_ticket_prices_pb2_grpc as pb2_grpc
from app.config import config
from app.celery import celery_app
import grpc



class GRPCClient:
    """Manages a gRPC connection pool."""
    def __init__(self):
        self.channel = grpc.insecure_channel(f'{config.PRIVATE_GRPC_HOST}:{config.PRIVATE_GRPC_PORT}')
        self.stub = pb2_grpc.PrivatePremierLeagueBookingStub(self.channel)

    def update_prices(self, match_id):
        request = pb2.UpdateTicketPricesRequest(match_id=match_id)
        print("I AM FUCKING REQUEST", request)
        return self.stub.UpdateTicketPrices(request)

# Create a single instance of the client
grpc_client = GRPCClient()

@celery_app.task
def update_ticket_prices(match_id: int):
    """Task to update ticket prices via gRPC."""
    try:
        response = grpc_client.update_prices(match_id)
        return {"status": response.status, "message": response.message}

    except Exception as e:
        return {"error": str(e)}
