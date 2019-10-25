# Bad request error message

An error message indicating that the subscription message was invalid.

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
			<td>400</td>
			<td>Identifier for a bad request error message</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Text</td>
			<td>String</td>
			<td>BADREQUEST</td>
			<td>The text of the message</td>
		</tr>
	</tbody>
</table>

> `400~BADREQUEST|`

# Bad API Key error message

An error message indicating that the API key is invalid.

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
			<td>401</td>
			<td>Identifier for a bad API key error message</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Text</td>
			<td>String</td>
			<td>BADAPIKEY</td>
			<td>The text of the message</td>
		</tr>
	</tbody>
</table>

> `401~BADAPIKEY|`

# Market not found error message

An error message indicating that the exchange and/or instrument in a subscription message was not found.

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
			<td>404</td>
			<td>Identifier for a market not found error message</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Text</td>
			<td>String</td>
			<td>NOTFOUND</td>
			<td>The text of the message</td>
		</tr>
	</tbody>
</table>

> `404~NOTFOUND|`

# Upstream not found error message

An error message indicating that the request couldn't be handled at this time, please try again in a few seconds.

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
			<td>503</td>
			<td>Identifier for an upstream not found error message</td>
		</tr>
		<tr>
			<td>Default separator</td>
			<td>Char</td>
			<td>~</td>
			<td>Default separator</td>
		</tr>
		<tr>
			<td>Text</td>
			<td>String</td>
			<td>UPSTREAMNOTFOUND</td>
			<td>The text of the message</td>
		</tr>
	</tbody>
</table>

> `503~UPSTREAMNOTFOUND|`