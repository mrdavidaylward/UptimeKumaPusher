# Uptime Kuma Push - Endpoint Testing 

This Go application is designed to test the availability of different types of endpoints, specifically ICMP, HTTP/S, and TCP endpoints, and send the result to the Uptime Kuma Push endpoint. It allows you to specify an interval for the checks to be performed continuously.

## Features

- Supports ICMP, HTTP/S, and TCP endpoint checks.
- Allows setting custom intervals between checks.
- Sends the status, message, and response time to the specified Uptime Kuma Push endpoint.

## Requirements

- Go 1.16 or later.

## Usage

### Flags

The application supports the following flags:

- `-apikey`: Your API Key (required)
- `-url`: URL to ping (required)
- `-kuma`: Uptime Kuma API URL (required)
- `-interval`: Interval between successive pings, in seconds (optional, default: 60)
- `-protocol`: Protocol for the test (icmp, http, https, tcp) (optional, default: icmp)
- `-port`: Port for TCP protocol (optional, default: 80)

### Example

You can run the application using the following command:

```bash
go run . -apikey=YOUR_PUSH_KEY -url=URL_TO_PING -kuma=UPTIME_API_URL
```

You can also set the protocol, interval, and port:


```bash
go run . -apikey=YOUR_PUSH_KEY -url=URL_TO_PING -kuma=UPTIME_KUMA_URL -protocol=http -interval=30 -port=8080
```
**Notes**

Please make sure that the server running this application has the necessary permissions to send ICMP packets if you plan on using ICMP for testing.

Remember to replace `YOUR_PUSH_KEY`, `URL_TO_PING`, and `UPTIME_KUMA_URL` with your actual values.
