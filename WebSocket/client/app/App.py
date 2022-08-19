import logging as log
import time

from structures.consts import Consts
from structures.context import Context
from structures.pygameApp import PygameApp
from structures.websocketApp import WebSocketApp


class App:
    def __init__(self, context: Context, consts: Consts):

        connecting = True
        while connecting:
            try:
                self.ws = WebSocketApp(context)
                connecting = False
            except ConnectionRefusedError or ConnectionAbortedError:
                log.warning('cannot connect')
                time.sleep(1)

        self.pygameApp = PygameApp(consts)
        self.context = context
        self.consts = consts

        self.ws.send_key()

    def one_iter(self):
        events = self.pygameApp.check_events()
        if 'quit' in events and events['quit']:
            # надо валить
            self.close()
            return

        self.ws.recv_msg()
        bodies = self.ws.get_bodies()
        self.ws.send_key()

        # типа сначала получили, потом нарисовали, а потом когда вернемя то уже будут тела (хз полезно ли)

        if not bodies:
            log.warning('bruh, no bodies')
            return

        self.pygameApp.draw_bodies(bodies)

    def run(self):
        while True:
            self.one_iter()
            time.sleep(1 / self.consts.frequency)

    def close(self):
        self.ws.close()
        self.pygameApp.close()
        quit()
        log.info('quiting')
