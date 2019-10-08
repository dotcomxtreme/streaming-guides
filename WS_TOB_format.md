# Level 1 (top of book) orderbook channel

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