from locust import HttpLocust, TaskSet, task

class CsdBehavior(TaskSet):

    @task
    def index(self):
        self.client.get("/")


class WebsiteUser(HttpLocust):
    task_set = CsdBehavior
    min_wait = 5000
    max_wait = 9000
    host = "http://api.istio.gcp.moncadaisla.es/ping"