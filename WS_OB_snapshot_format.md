# Orderbook Level 2 Snapshot Data Format

Snapshots are provided upon subscription in a raw ASCII format, as defined in the table below.

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

> `9~kraken~ETH~USD:40:208.98~27,208.27~1.303:208.99~13,209.12~1.1|`