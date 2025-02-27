import logging
import os
from dotenv import load_dotenv

# Load environment variables from .env

# if not succeeded:
#     logging.info(" load config successfully")
load_dotenv("config.env")

class Config:
    """Centralized configuration loader."""
    REDIS_HOST = os.getenv('REDIS_HOST')  # Default to 'localhost' if not set
    REDIS_PORT = os.getenv('REDIS_PORT')  # Default to '6379' if not set
    REDIS_DB = os.getenv('REDIS_DB')  # Default to '0' if not set

    # private grpc 
    PRIVATE_GRPC_HOST = os.getenv('PRIVATE_GRPC_HOST')
    PRIVATE_GRPC_PORT = os.getenv('PRIVATE_GRPC_PORT')

config = Config()  # Singleton instance
