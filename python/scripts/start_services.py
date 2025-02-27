import subprocess

def start_celery_worker():
    """Function to start the Celery worker."""
    worker_command = ['celery', '-A', 'app.celery', 'worker', '--loglevel=info']
    return subprocess.Popen(worker_command)

def start_celery_beat():
    """Function to start the Celery Beat."""
    beat_command = ['celery', '-A', 'app.celery_app', 'beat', '--loglevel=info']
    return subprocess.Popen(beat_command)
