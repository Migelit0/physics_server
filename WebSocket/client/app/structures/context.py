class Context:
    def __init__(self, server_ip: str, server_port: str, endpoint: str):
        self.server_ip = server_ip
        self.server_port = server_port
        self.url = f'ws://{server_ip}:{server_port}{endpoint}'
