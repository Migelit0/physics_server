import ast
import time

import websocket
import websocket as ws
import logging as log
import typing as tp
import rel

from context import Context
from body import Body


def on_error(ws, error):
    log.error('error: ', error)


def on_message(ws, msg):
    log.info(msg)


def on_close(ws, close_status_code, close_msg):
    log.info('### closed ###')
    log.info(close_status_code, close_msg)


def on_open(ws):
    log.info('Opened connection')


class WebSocketApp():  # ws.WebSocketApp):
    def __init__(self, context: Context):
        self.bodies: tp.Optional[tp.List[Body]] = None
        self.ws = ws.create_connection(context.url)
        self.context = context

    def recv_msg(self) -> tp.NoReturn:
        log.info('receiving msg')
        time.sleep(0.001)
        msg = self.ws.recv()
        log.info(f'get msg: {msg}')

        if not msg:
            return {}

        try:
            raw = dict(ast.literal_eval(msg))
            # чзх этот ваш аст (не нравится евал)
        except ValueError:
            log.warning('data error: ', msg)
            print(msg)
            return

        # FIXME: защита от инвалидной даты
        bodies = []
        for i, coords in raw.items():
            body = Body(int(i), coords[0], coords[1])
            bodies.append(body)

        self.bodies = bodies

    def send_key(self):
        log.info(f'sending: {self.context.key}')
        self.ws.send(self.context.key)

    def get_bodies(self) -> tp.List[Body]:
        return self.bodies

    def close(self):
        self.ws.send()
        self.ws.close()


if __name__ == '__main__':
    ws.enableTrace(True)
    context = Context('127.0.0.1', '30000', '/v1', 'bruh')
    # ws = WebSocketApp(context)
    ws_1 = websocket.WebSocketApp('ws://127.0.0.1:30000/v1', None, on_open, on_message, on_error, on_close)
    ws_1.run_forever(dispatcher=rel)
    rel.signal(2, rel.abort)

    # rel.dispatch()

    ###
    while 1:
        # ws.send('bruh')
        ws_1.send('bruh')
        # print(ws.recv())
        time.sleep(1)
