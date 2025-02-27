# from scripts.start_services import start_celery_worker
# from utils.load_config import load_config
# import logging


# def main():
#     # Set up logging configuration
#     logging.basicConfig(level=logging.INFO)
    
#     # load configuration file
#     succeed = load_config("config.env")
#     if succeed:
#        logging.info(" load config successfully")

#     # Start Celery worker and Beat
#     worker_process = start_celery_worker()
#     logging.info(" celery Worker are running.")

# if __name__ == "__main__":
#     main()
# import grpc
# import pb.rpc_update_ticket_prices_pb2 as pb2
# import pb.rpc_update_ticket_prices_pb2_grpc as pb2_grpc

import logging
from scripts.start_services import start_celery_worker
def main():
  
    worker_process = start_celery_worker()
    logging.info(" celery Worker are running.")

if __name__ == "__main__":
    main()