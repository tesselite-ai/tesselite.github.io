import os

from locust import FastHttpUser, task, between


HOSTS = {}

os.environ.get("SERVER","python")

class TestServer(FastHttpUser):
    wait_time = between(0.1, 0.5)

    host = f"http://{os.environ.get('SERVER', 'web')}:{os.environ.get('PORT', 8080)}"

    @task
    def hello_world(self):
        self.client.get("/")

