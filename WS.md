
# Websocket API
The Websocket API provides real time data updates for a supplied exchange over a websocket connection.

Data is sent using an ASCII streamer protocol with individual messages defined as below. Transmitted packets may contain one or several individual messages.

> `part1~part2~part3|`

Fields within messages are delimited by a ~ character and multiple messages are delimited by a | character which also serves as an end of message indicator.

The streaming service sends a heartbeat message at 10 second intervals to all connected Websocket API clients. The format is as follows:

> `999~HEARTBEAT|`

Once a client connection has been established it is necessary for the connecting client to request at least one subscription in order to receive data.

Clients who connect but do not open a subscription will be disconnected after a short grace period.

NB: The message formats are case-sensitive, with the exception of _**\[exchange]**_ which is case insensitive.

## Subscription messages

Data subscriptions are made for a given exchange / coin pair combination with one or many being passed in a single JSON object bearing the structure below.
> {
> 
>     "action": **_ACTION_TYPE_**
> 
>     "subs": [**_SOURCE1_**, **_SOURCE2_** ...]
> 
>     "api_key": **_API_KEY_**
> 
> }

### Action
Action type refers to whether the message is for a subscribe or unsubscribe event. It has the following acceptable values:

<table>
	<thead>
		<tr>
			<th>action param</th>
			<th>Meaning</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>SubAdd</td>
			<td>Subscribe to the included sources</td>
		</tr>
		<tr>
			<td>SubRemove</td>
			<td>Unsubscribe from the included sources</td>
		</tr>
	</tbody>
</table>

### Subs

The subs param is an array of strings containing qualified sources of the form:

> `[sourceid]~[exchange]~[coinfrom]~[cointo]`
> 
> or
> 
> `[sourceid]~[exchange]`

At least one source is required for a message to be valid.
Here are some potential examples.

<table>
	<thead>
		<tr>
			<th>Qualified source string</th>
			<th>Subscription result</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>30~bitfinex~BTC~USD</td>
			<td>Level 1 order book data for BTC to USD data for the Bitfinex exchange</td>
		</tr>
		<tr>
			<td>8~coinbase~BTC~USD</td>
			<td>Level 2 orderbook data for BTC to USD data for the Bitfinex exchange</td>
		</tr>
		<tr>
			<td>0~coinbase</td>
			<td>Trade data for all instruments on the Coinbase exchange</td>
		</tr>
	</tbody>
</table>

Successful subscriptions/un-subscriptions will be echoed back from the client.

Please refer to the sources field table below.
### Subscription sources

<table>
	<thead>
		<tr>
			<th>Source id</th>
			<th>Channel</th>
			<th>Entire Exchange Data Set Subscription Supported</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>0</td>
			<td>Trade data</td>
			<td>Yes</td>
		</tr>
		<tr>
			<td>30</td>
			<td>Level 1 (top of book) data</td>
			<td>No</td>
		</tr>
		<tr>
			<td>8</td>
			<td>Level 2 (aggregated) data</td>
			<td>No</td>
		</tr>
	</tbody>
</table>

## Trade channel

The message schema for trade updates is defined in the [Trade message format](WS_trade_format.md).

## Level 1 (top of book) orderbook channel

The top of book data will send messages of the format shown below whenever the best bid or best ask for a specified market changes. Alongside the price it will prove the current volume and effective exchange timestamp.

> `30~kraken~ETH~USD~1~10~100.10~6.05~1566398944495400100|`

The message schema for level 1 orderbook data is defined in [Orderbook top of book data format](WS_TOB_format.md).

On initial subscription the current best bid and best ask values will be received. Should one or both sides be missing, an [Orderbook top of book no data format](WS_TOB_no_value_format.md) message will be received.

## Level 2 orderbook channel

### Snapshot Data

When subscribing to level 2 orderbook channel an orderbook snapshot will be received during the initial collection messages. Each message has a sequence number and thus any messages with an earlier sequence number should be discarded.

Orderbook updates are categorised by a starting ID of 9 followed by the market information. An example is found below.

> `9~kraken~ETH~USD:40:208.98~27,208.27~1.303:208.99~13,209.12~1.1|`

The message schema for snapshots is defined in [Orderbook snapshot data format](WS_OB_snapshot_format.md).

### Incremental updates

After the snapshot, incremental updates will be received as per the example below.

> `8~kraken~ETH~USD~1~4~40~208.27~1.303|`

The message schema for incremental updates is defined in [Orderbook live data format](WS_OB_live_format.md).


