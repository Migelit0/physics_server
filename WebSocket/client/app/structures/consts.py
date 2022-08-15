import json

import typing as tp


class Consts:
    def __init__(self, frequency: int, height: int, width: int,
                 bg_color: tp.List[int], body_color: tp.List[int], ball_size: int):
        self.frequency = frequency
        self.height = height
        self.width = width
        self.background_color = bg_color
        self.body_color = body_color
        self.ball_size = ball_size


def generate_consts(path: str) -> Consts:
    with open(path, 'r') as file:
        data = file.read()

    json_data = json.loads(data)
    consts = Consts(json_data['frequency'],
                    json_data['height'],
                    json_data['width'],
                    json_data['background_color'],
                    json_data['body_color'],
                    json_data['ball_size'])

    return consts
