# GoLang ICMP Ping Endpoint Tester

This application is designed to test the availability of a specific endpoint using the ICMP ping command and then send the test results to the push API of Uptime Kuma.

## Requirements

- GoLang 1.14 or higher

## Usage

1. Replace YOUR_API_KEY with the API key.
2. Replace URL_TO_PING with the URL you want to ping.
3. Replace YOUR_UPTIME_KUMA_URL with your Uptime Kuma instance URL.
4. Set an interval that matched your uptime kuma interval.

For example:

```bash
go run main.go -apikey abc123 -url www.google.com -kuma https://your-uptime-kuma-instance.com
```
Flags
- apikey: Your API key.
- url: The URL of the endpoint you want to ping.
- kuma: The base URL of your Uptime Kuma instance.
- interval: The rate at which the polling rate of the endpoint, Default is 60s **should match what is set in uptime kuma
Functionality
The program pings the specified URL.
It checks the ping output and sets the status to "up" if any packets are received; otherwise, the status is set to "down".
The packet loss is extracted from the ping output.
The round-trip time (RTT) is extracted from the ping output.
The status, packet loss, and RTT are sent to the push API of Uptime Kuma.
Troubleshooting

**Ensure that you have the correct permissions to execute the ICMP ping command.
