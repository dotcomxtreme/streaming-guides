# Streamer API

## Document history

<table>
	<thead> 
		<tr>
			<th>Version</th>
			<th>Date</th>
			<th>Comments</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>1</td>
			<td>09/07/2019</td>
			<td>First public version</td>
		</tr>
		<tr>
			<td>1.1</td>
			<td>15/07/2019</td>
			<td>
                * Updated Top of book Orderbook Data Format to reflect changes in field order<br>
                * Added level 1 snapshot endpoint information<br>
                * minor tweaks
			</td>
		</tr>
		<tr>
			<td>1.2</td>
			<td>21/08/2019</td>
			<td>Added position and timestamp fields to live level 1 OB data</td>
		</tr>
		<tr>
			<td>1.3</td>
			<td>27/08/2019</td>
			<td>Added information about new subscription mechanism and clarified functionality</td>
		</tr>
		<tr>
			<td>1.4</td>
			<td>25/09/2019</td>
			<td>Updated to reflect orderbook snapshot received when subscribing to level 2 data</td>
		</tr>
	</tbody>
</table>

## Introduction
CryptoCompare offers a number of different streaming and polling endpoints for requesting and subscribing to market data.
This document serves as a developer guide for integrating with the market data streamer service for direct access to data via REST and WebSocket endpoints.
Please note this service is presently in a closed beta phase and subject to change.
The public interface is presently accessible at [https://interop.cryptocompare.com/](https://interop.cryptocompare.com/)

## REST API
The REST API provides a number of data endpoints as documented below.

### Orderbook snapshots
Snapshots can be requested for a given exchange and coin pair using the following URL:

> https://interop.cryptocompare.com/exchange-snapshot/api/getOrderBookByExchange/\<exchange>/\<coinfrom>-\<cointo>

The URL is case-sensitive, with the exception of _**\<exchange>**_ which is case insensitive.
For example, the streaming service can be queried for Bitfinex snapshots with instrument Bitcoin/USD as follows:

> https://interop.cryptocompare.com/exchange-snapshot/api/getOrderBookByExchange/bitfinex/BTC-USD

Snapshots are provided in a raw ASCII format, similar to the streaming protocol (discussed later), as defined in the table below.

<table>
	<thead>
		<tr>
			<th>Entry</th>
			<th>Type</th>
			<th>Possible values</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Message Type</td>
			<td>Integer</td>
			<td>9</td>
			<td>Identifier for a Full Orderbook</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange name</td>
			<td>String</td>
			<td>Alphanumeric (lowercase)</td>
			<td>The name of the Exchange to supply data for</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'from' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert from</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'to' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert to</td>
		</tr>
		<tr>
			<td>Header separator</td>
			<td>Char</td>
			<td>|</td>
			<td>The symbol denoting the beginning of the data payload</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>Integer</td>
			<td>1-65535</td>
			<td>Message sequence number as provided by the upstream data source</td>
		</tr>
		<tr>
			<td>Section separator</td>
			<td>Char</td>
			<td>:</td>
			<td>Separator for data sections</td>
		</tr>
		<tr>
			<td colspan=4>0 or more instances of the following 4 rows, sorted numerically on Bid Price descending</td>
		</tr>
		<tr>
			<td>Bid Price</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Bid Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Bid Quantity</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Bid Quantity</td>
		</tr>
		<tr>
			<td>Data separator</td>
			<td>Char</td>
			<td>,</td>
			<td>Separator for each price/quantity pair (omitted if this is the last pair)</td>
		</tr>
		<tr>
			<td>Section separator</td>
			<td>Char</td>
			<td>:</td>
			<td>Separator for data sections</td>
		</tr>
		<tr>
			<td colspan=4>0 or more instances of the following 4 rows, sorted numerically on Ask Price ascending</td>
		</tr>
		<tr>
			<td>Ask Price</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Ask Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Ask Quantity</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Ask Quantity</td>
		</tr>
		<tr>
			<td>Data separator</td>
			<td>Char</td>
			<td>,</td>
			<td>Separator for each price/quantity pair (omitted if this is the last pair)</td>
		</tr>
	</tbody>
</table>

For example, a snapshot for Kraken Ethereum to USD, containing 2 bids and 2 asks, and sent after the 40<sup>th</sup> Kraken ETH-USD update message was sent by the streaming service, may take the form:

> `9~kraken~ETH~USD|40:208.98~27,208.27~1.303:208.99~13,209.12~1.1`

### All Instruments

> https://interop.cryptocompare.com/exchange-snapshot/api/getAllKnownInstruments/

Will return a JSON object specifying all presently known instruments to this instance of the streaming service.

### Top of book

> https://interop.cryptocompare.com/exchange-snapshot/api/getTopOfOrderBooks/

Will return a JSON object specifying the best bid and best ask for all known instruments on this instance of the streaming service.

> https://interop.cryptocompare.com/exchange-snapshot/api/getTopOfOrderBooksByExchange/\<exchange>

Will return a JSON object specifying the best bid and best ask for all known instruments for a particular exchange.  

## Websocket API
The Websocket API provides real time data updates for a supplied exchange over a websocket connection.
Connections should be made to port 8080.
Data is sent using an ASCII streamer protocol with individual fields within messages are delimited by a ~ character and multiple messages are delimited by a | character.
The streaming service sends a heartbeat message at 10 second intervals to all connected Websocket API clients.  The format is as follows:

> `999~HEARTBEAT|`

Once a connection is established, it is necessary to send at least one subscription request to the streaming service in order to receive data. Connecting clients that do not send a subscription message will be disconnected after a short grace period.
NB: The message formats are case-sensitive, with the exception of _**\<exchange>**_ which is case insensitive.

### Subscription messages

There are two mechanisms for starting subscriptions; both are documented below. However it should be noted that the streaming message format for subscriptions is now deprecated and will be removed in the future.

#### **JSON streamer message**

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

**Please note:** A valid API key is required for this subscription message type

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
			<td>Unsubscribe to the included sources</td>
		</tr>
	</tbody>
</table>

The subs param is an array of strings containing qualified sources of the form:

> `<sourceid>~<exchange>~<coinfrom>~<cointo>`
> 
> or
> 
> `<sourceid>~<exchange>`

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

#### **Streamer message (legacy)**

> **WARNING:** This subscription type is no longer supported and will likely be removed with little warning in the future

Data subscriptions are made for a given exchange / coin pair combination.  The message format is as follows:

> `SUBSCRIBE~<sourceid>~<exchange>~<coinfrom>~<cointo>|`
>
> or
>
> `SUBSCRIBE~<sourceid>~<exchange>|`

The type of subscription that will be opened will depend on the supplied _sourceid_ in the message above. Where no _coinfrom_ and _cointo_ params are supplied, a subscription will be opened for the entire exchange data set if the _sourceid_ supports this.
Multiple messages are required to subscribe to multiple exchange/instrument/sources and can be sent on connection start or any time thereafter. Unsubscribing from sources is performed via the UNSUBSCRIBE command and follows the same pattern.

> `UNSUBSCRIBE~<sourceid>~<exchange>~<coinfrom>~<cointo>|`
> 
> or
> 
> `UNSUBSCRIBE~<sourceid>~<exchange>|`

Successful subscriptions/un-subscriptions will be echoed back from the client.

##### For example, to subscribe to live Bitcoin to USD data for the Bitfinex exchange, the following message should be sent via the websocket:

> `SUBSCRIBE~8~bitfinex~BTC~USD|`

Subscriptions for level 1 (best bid / best ask) are made in the same way (with a given exchange / coin pair combination specified) however with a different source ID.

##### For example, to subscribe to live Bitcoin to USD data for the Bitfinex exchange, the following message should be sent via the websocket:
> `SUBSCRIBE~30~bitfinex~BTC~USD|`

Please refer to the sources field table below.
#### Subscription sources

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

### Orderbook Data Format
#### Snapshot Data
Snapshots are provided upon subscription in a raw ASCII format, as defined in the table below.

<table><colgroup><col><col><col><col></colgroup>
    <thead>
		<tr>
			<th>Entry</th>
			<th>Type</th>
			<th>Possible values</th>
			<th>Description</th>
		</tr>
    </thead>
	<tbody>
		<tr>
			<td>Message Type</td>
			<td>Integer</td>
			<td>9</td>
			<td>Identifier for a Full Orderbook</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange name</td>
			<td>String</td>
			<td>Alphanumeric (lowercase)</td>
			<td>The name of the Exchange to supply data for</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'from' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert from</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'to' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert to</td>
		</tr>
		<tr>
			<td>Header separator</td>
			<td>Char</td>
			<td>~</td>
			<td>The symbol denoting the beginning of the data payload</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>Integer</td>
			<td>1-65535</td>
			<td>Message sequence number as provided by the upstream data source</td>
		</tr>
		<tr>
			<td>Section separator</td>
			<td>Char</td>
			<td>:</td>
			<td>Separator for data sections</td>
		</tr>
		<tr>
			<td colspan=4>0 or more instances of the following 4 rows, sorted numerically on Bid Price descending</td>
		</tr>
		<tr>
			<td>Bid Price</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Bid Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Bid Quantity</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Bid Quantity</td>
		</tr>
		<tr>
			<td>Data separator</td>
			<td>Char</td>
			<td>,</td>
			<td>Separator for each price/quantity pair (omitted if this is the last pair)</td>
		</tr>
		<tr>
			<td>Section separator</td>
			<td>Char</td>
			<td>:</td>
			<td>Separator for data sections</td>
		</tr>
		<tr>
			<td colspan=4>0 or more instances of the following 4 rows, sorted numerically on Ask Price ascending</td>
		</tr>
		<tr>
			<td>Ask Price</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Ask Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Ask Quantity</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Ask Quantity</td>
		</tr>
		<tr>
			<td>Data separator</td>
			<td>Char</td>
			<td>,</td>
			<td>Separator for each price/quantity pair (omitted if this is the last pair)</td>
		</tr>
	</tbody>
</table>

For example, a snapshot for Kraken Ethereum to USD, containing 2 bids and 2 asks, and sent after the 40<sup>th</sup> Kraken ETH-USD update message was sent by the streaming service, may take the form:

> `9~kraken~ETH~USD:40:208.98~27,208.27~1.303:208.99~13,209.12~1.1`

#### Live Data

Live Orderbook data will be sent in the format shown in the table below.

<table>
	<thead>
		<tr>
			<th>Entry</th>
			<th>Type</th>
			<th>Possible values</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Message Type</td>
			<td>Integer</td>
			<td>8</td>
			<td>Identifier for an Orderbook update</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange name</td>
			<td>String</td>
			<td>Alphanumeric (lowercase)</td>
			<td>The name of the Exchange to supply data for</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'from' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert from</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'to' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert to</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Side</td>
			<td>Integer</td>
			<td>
                1 – Bid<br>
                2 – Ask
			</td>
			<td>Indicates if the message represents a buy or sell request</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Type Flag</td>
			<td>Integer</td>
			<td>
                1 – Add<br>
                2 – Remove<br>
                4 – Change/update<br>
			</td>
			<td>Describes the type of message sent</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>Integer</td>
			<td>1-65535</td>
			<td>Message sequence number as provided by the upstream data source</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Price</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Quantity</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Quantity</td>
		</tr>
		<tr>
			<td>Message separator</td>
			<td>Char</td>
			<td>|</td>
			<td>The symbol denoting the end of the data payload</td>
		</tr>
	</tbody>
</table>

For example, a bid update for Kraken Ethereum to USD, which is flagged by the upstream data source as the 40th Kraken ETH-USD message sent, may take the form:

> `8~kraken~ETH~USD~1~4~40~208.27~1.303|`

### Top of book Orderbook Data Format
Live Orderbook data will be sent in the format shown in the table below.

<table>
	<thead>
		<tr>
			<th>Entry</th>
			<th>Type</th>
			<th>Possible values</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Message Type</td>
			<td>Integer</td>
			<td>30</td>
			<td>Identifier for a level 1 Orderbook update</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange name</td>
			<td>String</td>
			<td>Alphanumeric (lowercase)</td>
			<td>The name of the Exchange to supply data for</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'from' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert from</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'to' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert to</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Side</td>
			<td>Integer</td>
			<td>
                1 – Bid<br>
                2 – Ask<br>
			</td>
			<td>Indicates if the following price represents a best bid or a best ask</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>Integer</td>
			<td>1-65535</td>
			<td>Message sequence number as provided by the upstream data source</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Price</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Quantity</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Quantity</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Timestamp</td>
			<td>Integer</td>
			<td>Signed 64 bit nanoseconds</td>
			<td>Last time the price was updated, as reported by the upstream data source in nanoseconds (Units: Unix time * 10^9)</td>
		</tr>
		<tr>
			<td>Message separator</td>
			<td>Char</td>
			<td>|</td>
			<td>The symbol denoting the end of the data payload</td>
		</tr>
	</tbody>
</table>

For example, a best bid field update for Kraken Ethereum to USD, which is flagged by the upstream data source as the 10th Kraken ETH-USD message sent, may take the form:

> `30~kraken~ETH~USD~1~10~100.10~6.05~1566398944495400100|`

### Trade Data Format
Live Trade data will be sent in the format shown in the table below.

<table>
	<thead>
		<tr>
			<th>Entry</th>
			<th>Type</th>
			<th>Possible values</th>
			<th>Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Message Type</td>
			<td>Integer</td>
			<td>0</td>
			<td>Identifier for a Trade update</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange name</td>
			<td>String</td>
			<td>Alphanumeric (lowercase)</td>
			<td>The name of the Exchange to supply data for</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'from' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert from</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Coin 'to' symbol</td>
			<td>String</td>
			<td>Alphanumeric (uppercase)</td>
			<td>The currency to convert to</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Side</td>
			<td>Integer</td>
			<td>
            1 – Buy<br>
            2 – Sell<br>
            4 – Unknown<br>
			</td>
			<td>Indicates if the message represents a buy or sell request</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Exchange ID</td>
			<td>String</td>
			<td>ASCII</td>
			<td>Trade ID as provided by the Exchange</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Timestamp</td>
			<td>Integer</td>
			<td>Signed 64 bit microseconds</td>
			<td>Time of trade as reported by the Exchange (Units: Unix time * 10^6)</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Quantity</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Last Volume</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Price</td>
			<td>String</td>
			<td>0.0+</td>
			<td>Last Price</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Total</td>
			<td>Float</td>
			<td>0.0+</td>
			<td>Quantity multiplied by Price - i.e. the total cost of the transaction</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>ID</td>
			<td>Integer</td>
			<td>Signed 64 bit nanoseconds</td>
			<td>Unique Identifier: Time of receipt of trade message by the upstream data source in nanoseconds (Units: Unix time * 10^9)</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>Integer</td>
			<td>1-65535</td>
			<td>Message sequence number as provided by the upstream data source</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Presence Flag for "TradeFields" Entries</td>
			<td>String</td>
			<td>Two digit lowercase hex</td>
			<td>Bitflag used to indicate which of the entries defined in the "Streamer TradeFields Bitflags" table below have been included. Intended for futureproofing - at present, none of the entries are optional, with the result that the bitflag will be the same in every message.</td>
		</tr>
		<tr>
			<td>Message separator</td>
			<td>Char</td>
			<td>|</td>
			<td>The symbol denoting the end of the data payload</td>
		</tr>
	</tbody>
</table>


<table>
	<thead>
		<tr>
			<th colspan=2>Streamer TradeFields Bitflags</th>
		</tr>
		<tr>
			<th>Entry</th>
			<th>Flag</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Exchange ID</td>
			<td>1 (0x01)</td>
		</tr>
		<tr>
			<td>Timestamp</td>
			<td>10 (0x02)</td>
		</tr>
		<tr>
			<td>Quantity</td>
			<td>100 (0x04)</td>
		</tr>
		<tr>
			<td>Price</td>
			<td>1000 (0x08)</td>
		</tr>
		<tr>
			<td>Total</td>
			<td>10000 (0x10)</td>
		</tr>
		<tr>
			<td>ID</td>
			<td>100000 (0x20)</td>
		</tr>
		<tr>
			<td>Sequence number</td>
			<td>1000000 (0x40)</td>
		</tr>
	</tbody>
</table>

For example, a "Sell" trade update for Kraken Ethereum to USD, which is flagged by the upstream data source as the 50th Kraken trade message sent, received on Friday, 30-Nov-18 at 17:30:31 UTC, and for which the Exchange provided a Trade ID of "TRADE-ABCDEFGHIJ", which the Exchange reported traded at 17:30:30, may take the form:

> `0~kraken~ETH~USD~2~TRADE-ABCDEFGHIJ~1543599030123456~50~217.35~2.348~1543599031123456789~7f|`