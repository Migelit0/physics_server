import ssl

import websocket as ws


def on_message(ws, message):
    print(ws)
    print(message)


def on_error(ws, error):
    print(ws)
    print(error)


def on_close(ws):
    print(ws)


# if __name__ == '__main__':
def main():
    sslopt = {"cert_reqs": ssl.CERT_NONE, "check_hostname": False}
    conn = ws.create_connection('wss://127.0.0.1:993/test/', sslopt=sslopt)

    print('testing conn')
    conn.send('test')
    print(conn.recv())


if __name__ == '__main__':
    main()
