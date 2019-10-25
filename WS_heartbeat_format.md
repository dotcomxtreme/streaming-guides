# Heartbeat format

A heartbeat message sent every 10 seconds.

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
			<td>999</td>
			<td>Identifier for a heartbeat</td>
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
			<td>HEARTBEAT</td>
			<td>The text of the message</td>
		</tr>
	</tbody>
</table>

> `999~HEARTBEAT|`
