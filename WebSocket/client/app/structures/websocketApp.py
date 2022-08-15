import time

import websocket as ws
import logging as log
import typing as tp
import rel

from context import Context
from body import Body


def on_error(ws, error):
    log.error('error: ', error)


def on_close(ws, close_status_code, close_msg):
    log.info('### closed ###')
    log.info(close_status_code, close_msg)


def on_open(ws):
    log.info('Opened connection')


class WebSocketApp(ws.WebSocketApp):
    def __init__(self, context: Context, header=None, on_open=on_open, on_error=on_error,
                 on_close=on_close, on_ping=None, on_pong=None, on_cont_message=None, keep_running=True,
                 get_mask_key=None, cookie=None, subprotocols=None, on_data=None):
        url: str = context.url
        on_message = self.on_message
        self.bodies: tp.Optional[tp.List[Body]] = None

        super().__init__(url, header, on_open, on_message, on_error, on_close, on_ping, on_pong, on_cont_message,
                         keep_running, get_mask_key, cookie, subprotocols, on_data)
        self.context = context

    def on_message(self, ws, message) -> tp.NoReturn:
        log.info('get msg: ', message)
        raw = dict(message)
        bodies = []
        for i, coords in raw.items():
            body = Body(int(i), coords[0], coords[1])
            bodies.append(body)

        self.bodies = bodies

    def get_bodies(self) -> tp.List[Body]:
        return self.bodies


if __name__ == '__main__':
    context = Context('127.0.0.1', '30000', '/v1')
    ws = WebSocketApp(context)
    ws.run_forever(dispatcher=rel)
    while 1:
        ws.send('bruh')
        # print(ws.recv())
        time.sleep(1)
