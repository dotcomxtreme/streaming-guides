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
			<td>08/10/2019</td>
			<td>First public version</td>
		</tr>
	</tbody>
</table>

## Introduction
CryptoCompare offers a number of different streaming and polling endpoints for requesting and subscribing to market data. REST requests may be performed via our [polling API](https://min-api.cryptocompare.com/). Streaming via websocket provides an alternative way of consuming our data in a reliable, low latency manner. 

This document serves as a developer guide for integrating with the market data streamer service.

Please note this service is presently in closed beta and subject to change.
The public interface is presently accessible via [wss://streaming.cryptocompare.com/](wss://streaming.cryptocompare.com/) and will require IP whitelisting as well as an API key. This can be requested by contacting us [by email.](mailto:data@cryptocompare.com)

### Examples

<table>
	<thead> 
		<tr>
			<th>Language</th>
			<th>Samples</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Python 3</td>
			<td>
				<ul>
					<li><a href="examples/python/level-1-streaming.py">Level 1 streaming</a></li>
					<li><a href="examples/python/level-2-streaming.py">Level 2 streaming</a></li>
				</ul>
			</td>
		</tr>
		<tr>
			<td>NodeJS</td>
			<td>
				<ul>
					<li><a href="examples/nodejs/level-1-streaming.js">Level 1 streaming</a></li>
					<li><a href="examples/nodejs/level-2-streaming.js">Level 2 streaming</a></li>
				</ul>
			</td>
		</tr>
		<tr>
			<td>Golang</td>
			<td>
				<ul>
					<li><a href="examples/golang/streaming-example.go">Trade, level 1 and 2 streaming</a></li>
				</ul>
			</td>
		</tr>
	</tbody>
</table>

Full implementation details can be found in the [Websocket API reference](WS.md).
