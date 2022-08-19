import sys
import os
import logging as log

log.basicConfig(level=log.INFO)

sys.path.append('app')
sys.path.append('app/structures')

from app.structures.consts import generate_consts
from app.structures.context import generate_context
from app.App import App


def main():
    context_path = os.path.abspath('secrets/context.json')
    consts_path = os.path.abspath('secrets/consts.json')

    context = generate_context(context_path)
    consts = generate_consts(consts_path)

    app = App(context, consts)
    app.run()


if __name__ == '__main__':
    main()
