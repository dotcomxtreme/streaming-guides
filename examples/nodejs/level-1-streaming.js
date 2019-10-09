const ws = require('ws');

function parseLevel1(message) {
    let dataArray = message.split('~');
    let dataObj = {
        type: dataArray[0],
        market: dataArray[1],
        fromsymbol: dataArray[2],
        tosymbol: dataArray[3],
        side: dataArray[4] === '1' ? 'bid' : 'ask',
        sequence: parseInt(dataArray[5], 10),
        price: parseFloat(dataArray[6]),
        quantity: parseFloat(dataArray[7]),
        timestamp: new Date(parseInt(dataArray[8],10)/1000000)
    }
    return dataObj;
}


let cryptocompareWS = new ws('wss://streaming.cryptocompare.com');

let subs = {
    "action": "SubAdd",
    "subs": ["30~bitfinex~BTC~USD"],
    "api_key": "YOUR-API-KEY",
    "format": "streamer"
}

cryptocompareWS.on('open', function(){
    console.log('connection established')
    cryptocompareWS.send(JSON.stringify(subs));
});

cryptocompareWS.on('error', function(error){
    console.log(error)
});

cryptocompareWS.on('close', function(){
    console.log('disconnected');
});

cryptocompareWS.on('message', function(data){
    console.log(data)
    if (data === '999~HEARTBEAT|') {
        console.log(data);
        return;
    }
    let updates = data.split('|');
    updates.pop(); //Remove last element as it is always an empty string
    for (let update of updates) {
        let parsedUpdate = parseLevel1(update);
        console.log(JSON.stringify(parsedUpdate))
    }
});