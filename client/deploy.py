import json

import requests


class Deploy:

    def __init__(self, url: str):
        self.url = url
        self.data = None

    def get_states_dict(self):
        url = self.url + '/dict'
        response = requests.get(url)
        data = json.loads(response.text)
        print(json.dumps(data, indent=4))

    def upload(self):
        pass

    def push(self):
        pass

    def clean(self):
        pass

    def build(self):
        pass

    def status(self):
        url = self.url + '/status'
        response = requests.get(url)
        data = json.loads(response.text)
        print(json.dumps(data, indent=4))


def main():
    url = 'http://localhost:8000'
    deploy = Deploy(url)
    deploy.get_states_dict()
    deploy.status()


if __name__ == '__main__':
    main()
