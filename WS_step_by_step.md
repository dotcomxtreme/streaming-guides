# Step by step example
The following descrbes a language agnostic method for requesting what data is available using CryptoCompare REST API endpoints and then using this to stream data via the streaming service.

## Step 1 : Determining available markets

To understand what markets are available for streaming, first we need to call [https://min-api.cryptocompare.com/data/ob/l2/exchanges](https://min-api.cryptocompare.com/documentation?key=Orderbook&cat=exchangesWithOrdebookStaticInfoEndpoint) which may return the following.

```
{
    "Response":"Success",
    "Message":"",
    "HasWarning":false,
    "Type":100,
    "RateLimit":{},
    "Data":
    [
        "Binance",
        "BitBank",
        // ... etc
        "bitFlyer",
        "itBit"
    ]
}
```

From this we can tell what exchange venues are available on the streaming service. These can either be subscribed to directly for updates on all their markets or we can now make a follow up REST request to determine which markets are currently provided.

[https://min-api.cryptocompare.com/data/v3/all/exchanges?e=exchangeName](https://min-api.cryptocompare.com/data/v3/all/exchanges?e=exchangeName)


An example output for Coinbase might be

```
{
    "Response":"Success",
    "Message":"",
    "HasWarning":false,
    "Type":100,
    "RateLimit":{},
    "Data":
    {
        "Coinbase":
        {
            "pairs":
            {
                "ETH":["DAI","USD","USDC","EUR","GBP","BTC"],
                "LINK":["USD","ETH"],
                // ... etc
                "BCH":["BTC","GBP","EUR","USD"],
                "ZEC":["USDC","BTC"],
            },
            "isActive":true,
            "isTopTier":true
            }
        }
    }
}
```

We now know the available markets on Coinbase allowing us to subscribe to them directly via the streaming WS endpoint.

## Step 2: Subscribing to streaming data

As an example, let's assume we want to subscribe to streaming Coinbase trade and level 1 (top of book) data for ETH-USD and ETH-BTC. On opening a websocket connection to wss://streaming.cryptocompare.com, we would send the following subscription message.

```
{
    "action": "SubAdd"
    "subs": 
    [
        "0~coinbase~ETH~USD",
        "30~coinbase~ETH~USD",
        "0~coinbase~ETH~BTC",
        "30~coinbase~ETH~BTC"
    ]
    "api_key": "YOUR_API_KEY"
}
```

Alongside heartbeat messages, the streaming service should respond with an echo of all opened subscriptions.

```
Received  {"subs":["0~Coinbase~ETH~USD","30~Coinbase~ETH~USD","0~Coinbase~ETH~BTC","30~Coinbase~ETH~BTC"]}

Received  30~coinbase~ETH~USD~2~1842~189.48~9.21450063~1570705179758638070|
Received  30~coinbase~ETH~USD~1~1843~189.49~8.09940423~1570705180858827602|
Received  30~coinbase~ETH~USD~2~1844~189.5~9.21450063~1570705182945202474|
Received  30~coinbase~ETH~USD~2~1847~189.54~60.13133817~1570705185905734886|
...
```

Updates should begin streaming immediately.