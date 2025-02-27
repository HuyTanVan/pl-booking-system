import logging
import os
from celery import Celery

from app.config import config
# selery using redis as broker


redis_url = f'redis://{config.REDIS_HOST}:{config.REDIS_PORT}/{config.REDIS_DB}'

celery_app = Celery(
    'celery_app',
    broker=redis_url,  # Redis as the message broker
    backend=redis_url,  # Redis as the result backend
    include=['app.tasks']
)


# Optional configuration, see the application user guide.
celery_app.conf.update(
    result_expires=3600,
    worker_pool='solo',
    broker_connection_retry_on_startup = True,
)
# app.conf.broker_connection_retry_on_startup = True

if __name__ == '__main__':
    celery_app.start()