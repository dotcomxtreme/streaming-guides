const ws = require('ws');


function parseSnapshot(message) {
    let subSections = message.split(':');
    let messageKey = subSections[0];
    let bidsString = subSections[1];
    let asksString = subSections[2];
    
    let fullOB = {}
    fullOB.bids = bidsString.split(',').map(function(item){
        return item.split('~').map(parseFloat);
    });
    fullOB.asks = asksString.split(',').map(function(item){
        return item.split('~').map(parseFloat);
    });
    return fullOB;
}

function getChangeFlag(flag) {
    if (flag === '1') {
        return 'add';
    } else if (flag === '2') {
        return 'remove';
    } else if (flag === '4') {
        return 'update';
    }
    return 'unknown'; //This should not happen
}

function parseLevel2(message) {
    let dataArray = message.split('~');
    let dataObj = {
        type: dataArray[0],
        exchange: dataArray[1],
        fromsymbol: dataArray[2],
        tosymbol: dataArray[3],
        side: dataArray[4] === '1' ? 'bid' : 'ask',
        change: getChangeFlag(dataArray[5]),
        sequence: parseInt(dataArray[6]),
        price: parseFloat(dataArray[7]),
        quantity: parseFloat(dataArray[8])
    }
    return dataObj;
}

let cryptocompareWS = new ws('wss://streaming.cryptocompare.com')

let subs = {
    "action": "SubAdd",
    "subs": ["8~bitfinex~BTC~USD"],
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
    let messages = data.split('|');
    messages.pop(); //Remove last element as it is always an empty string
    for (let message of messages) {
        if (message.split('~')[0] === '9') {
            let snapshot = parseSnapshot(message);
            console.log(snapshot)
            continue;
        }
        let parsedUpdate = parseLevel2(message);
        
        console.log(JSON.stringify(parsedUpdate))
    }
});

