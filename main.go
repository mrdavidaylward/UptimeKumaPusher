package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	apiKey    = flag.String("apikey", "", "API Key")
	targetUrl = flag.String("url", "", "URL to ping")
	apiUrl    = flag.String("kuma", "", "uptime kuma url") // API base url
	interval  = flag.Int("interval", 60, "Interval between successive pings, in seconds")
	protocol  = flag.String("protocol", "icmp", "Protocol for the test (icmp, http, https, tcp)")
	port      = flag.String("port", "80", "Port for TCP protocol")
)

func main() {
	flag.Parse()



	if *apiKey == "" || *targetUrl == "" || *apiUrl == "" {
		flag.Usage()
		return
	}

	for {
		startTime := time.Now()

		var status string
		var msg string
		var pingTime string

		switch *protocol {
		case "icmp":
			status, msg, pingTime = testICMP()
		case "http", "https":
			status, msg, pingTime = testHTTP(*protocol)
		case "tcp":
			status, msg, pingTime = testTCP()
		default:
			fmt.Printf("Unsupported protocol: %s\n", *protocol)
			return
		  }

		sendPingData(status, msg, pingTime)
		pingDuration := time.Since(startTime)
		if *interval > int(pingDuration.Seconds()) {
			time.Sleep(time.Duration(*interval-int(pingDuration.Seconds())) * time.Second)
		}
	}
}
