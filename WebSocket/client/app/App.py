from structures.consts import Consts
from structures.context import Context
from structures.pygameApp import PygameApp
from structures.websocketApp import WebSocketApp


class App:
    def __init__(self, context: Context, consts: Consts):
        self.ws = WebSocketApp(context)
        self.pygameApp = PygameApp(consts)
        self.context = context
        self.consts = consts

    def one_iter(self):
        bodies = self.ws.get_bodies()
        self.pygameApp.draw_bodies(bodies)

    def run(self):
        while True:
            self.one_iter()
