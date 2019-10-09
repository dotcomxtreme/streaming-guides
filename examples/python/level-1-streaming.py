from lomond import WebSocket
from lomond.persist import persist

def parse_level_1(message):
    data_array = message.split('~');
    data_obj = {
        "message_type": data_array[0],
        "exchange": data_array[1],
        "fromsymbol": data_array[2],
        "tosymbol": data_array[3],
        "side": 'bid' if data_array[4] == '1' else 'ask',
        "sequence": int(data_array[5]),
        "price": float(data_array[6]),
        "quantity": float(data_array[7]),
        "timestamp": int(data_array[8])/1000000
    }
    return data_obj

cryptocompare_ws = WebSocket('wss://streaming.cryptocompare.com')
subs = {
    "action": "SubAdd",
    "subs": ["30~bitfinex~BTC~USD"],
    "api_key": "YOUR-API-KEY",
    "format": "streamer"
}

for event in persist(cryptocompare_ws):
    if event.name == "ready":
        print('ready')
        cryptocompare_ws.send_json(subs)
    if event.name == "text":
        messages = event.text.split("|")
        messages = messages[:-1] #Last element is empty string
        for message in messages:
            print(parse_level_1(message))