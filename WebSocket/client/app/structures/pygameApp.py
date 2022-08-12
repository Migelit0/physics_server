import typing as tp

import pygame

from body import Body
from consts import Consts


class App:
    def __init__(self, consts: Consts):
        self.consts = consts
        self.init_pygame()

    def init_pygame(self):
        pygame.init()
        size = self.consts.width, self.consts.height
        self.screen = pygame.display.set_mode(size)
        self.clock = pygame.time.Clock()

    def draw_bodies(self, bodies: tp.List[Body]):
        self.screen.fill(self.consts.background_color)
        for body in bodies:
            pygame.draw.circle(self.screen, self.consts.body_color, (body.x, body.y), self.consts.ball_size)
