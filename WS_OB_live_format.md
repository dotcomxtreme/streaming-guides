# Orderbook Level 2 Live Data Format

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
			<td>Exchange Timestamp (optional)</td>
			<td>String</td>
			<td>1570803710878860000</td>
			<td>Timestamp (expressed in nanoseconds) as provided by exchange, optional as not all exchanges provide this</td>
		</tr>
		<tr>
			<td>Presence bit flag</td>
			<td>String</td>
			<td>0-1</td>
			<td>Indicated which optional fields are present in packed message. Please refer to table below.</td>
		</tr>
		<tr>
			<td>Message separator</td>
			<td>Char</td>
			<td>|</td>
			<td>The symbol denoting the end of the data payload</td>
		</tr>
	</tbody>
</table>

## Presence bit flag

The bitflag will indicate which optional fields are present in the message by combining the bits into a hex formatted numeric ASCII value. Irrespective of which fields are included in a message, they will be ordered based on numeric sequence as defined below.

<table>
	<thead>
		<tr>
			<th>Entry</th>
			<th>Flag</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Exchange timestamp</td>
			<td>1 (0x01)</td>
		</tr>
	</tbody>
</table>


For example, a bid update for Kraken Ethereum to USD without an exchange timestamp, which is flagged by the upstream data source as the 40th Kraken ETH-USD message sent, may take the form:

> `8~kraken~ETH~USD~1~4~40~208.27~1.303~0|`

A bid update for Kraken Ethereum to USD with an exchange timestamp, which is flagged by the upstream data source as the 40th Kraken ETH-USD message sent, may take the form:

> `8~kraken~ETH~USD~1~4~40~208.27~1.303~1570803710878860000~1|`


