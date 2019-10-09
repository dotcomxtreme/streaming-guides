# Trade Data Format

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
			<td>Presence flag for entries</td>
			<td>String</td>
			<td>Two digit lowercase hex</td>
			<td>Bitflag used to indicate which of the entries defined in Presence bit flag table below have been included within this message</td>
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

The bitflag will indicate which fields are present in the message by combining the bits into a hex formatted numeric ASCII value. Irrespective of which fields are included in a message, they will be ordered based on numeric sequence as defined below.

<table>
	<thead>
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

A message containing all of these fields would result in a bit flag value of **7f**. This is arrived at by the combination of all values (0x01 - 0x40) resulting in the hex value 0x7f which is represented as a string value. If one or more fields are omitted this value will be different. 

> `0~EXCHANGE~FROMSYM~TOSYM~SIDE~EXCHANGEID~TIMESTAMP~QUANTITY~PRICE~TOTAL~ID~7f|`

In the below example, the trade message does not contain an Exchange ID (0x01) thus has a bitflag value of **7e** (0x7f-0x01).

> `0~EXCHANGE~FROMSYM~TOSYM~SIDE~TIMESTAMP~QUANTITY~PRICE~TOTAL~ID~7e|`