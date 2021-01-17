import json

import requests


def print_name(func):
    def wrapper(*args):
        print(func.__name__)
        func(*args)
    return wrapper

class Deploy:

    def __init__(self, url: str):
        self.url = url
        self.data = None

    @print_name
    def get_states_dict(self):
        url = self.url + '/dict'
        response = requests.get(url)
        data = json.loads(response.text)
        print(json.dumps(json.loads(response.text), indent=4))

    @print_name
    def upload(self, fname: str):
        url = self.url + '/upload'
        files = {'package': open(fname, 'rb')}
        response = requests.post(url, files=files)
        print(json.dumps(json.loads(response.text), indent=4))

    @print_name
    def push(self):
        pass

    @print_name
    def clean(self):
        pass

    @print_name
    def build(self):
        url = self.url + '/build'
        json_data= {'apps': ['app1', 'app2'], 'version': 'v1.0'}
        response = requests.post(url, json=json_data)
        print(json.dumps(json.loads(response.text), indent=4))

    @print_name
    def status(self):
        url = self.url + '/status'
        response = requests.get(url)
        print(json.dumps(json.loads(response.text), indent=4))


def main():
    url = 'http://localhost:8000'
    deploy = Deploy(url)
    deploy.get_states_dict()
    deploy.status()
    deploy.upload('./a.tar')
    deploy.status()
    deploy.build()
    deploy.status()


if __name__ == '__main__':
    main()
