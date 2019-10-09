from lomond import WebSocket
from lomond.persist import persist

def parse_snapshot(message):
    sub_sections = message.split(':');
    message_key = sub_sections[0];
    bids_str = sub_sections[1];
    asks_str = sub_sections[2];
    full_OB = {}
    full_OB["bids"] = [[float(item) for item in position.split('~')] for position in bids_str.split(',')]
    full_OB["asks"] = [[float(item) for item in position.split('~')]  for position in asks_str.split(',')]
    return full_OB;


def get_change_flag(flag): 
    if (flag == '1'):
        return 'add'
    elif (flag == '2'):
        return 'remove'
    elif (flag == '4'):
        return 'update'
    return 'unknown'; #This should not happen

def parse_level_2(message):
    data_array = message.split('~');
    data_obj = {
        "message_type": data_array[0],
        "exchange": data_array[1],
        "fromsymbol": data_array[2],
        "tosymbol": data_array[3],
        "side": 'bid' if data_array[4] == '1' else 'ask',
        "change": get_change_flag(data_array[5]),
        "sequence": int(data_array[6]),
        "price": float(data_array[7]),
        "quantity": float(data_array[8])
    }
    return data_obj

cryptocompare_ws = WebSocket('wss://streaming.cryptocompare.com')
subs = {
    "action": "SubAdd",
    "subs": ["8~bitfinex~BTC~USD"],
    "api_key": "YOUR-API-KEY",
    "format": "streamer"
}

for event in persist(cryptocompare_ws):
    if event.name == "ready":
        print('ready')
        cryptocompare_ws.send_json(subs)
    if event.name == "text":
        print(event.text)
        messages = event.text.split("|")
        messages = messages[:-1] #Last element is empty string
        for message in messages:
            if message == "999~HEARTBEAT|":
                print(message)
            elif message.split("~")[0] == "9":
                print(parse_snapshot(message))
                continue
            else:
                print(parse_level_2(message))

            