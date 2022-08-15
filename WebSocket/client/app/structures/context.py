import typing as tp

import json


class Context:
    def __init__(self, server_ip: str, server_port: str, endpoint: str):
        self.server_ip = server_ip
        self.server_port = server_port
        self.url = f'ws://{server_ip}:{server_port}{endpoint}'


def generate_context(path: str) -> Context:
    with open(path, 'r') as file:
        data = file.read()

    json_data = json.loads(data)
    cont = Context(json_data['server_ip'],
                   json_data['server_port'],
                   json_data['endpoint'])
    return cont
