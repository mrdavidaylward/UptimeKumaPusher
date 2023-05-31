package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	apiKey    = flag.String("apikey", "", "API Key")
	targetUrl = flag.String("url", "", "URL to ping")
	apiUrl    = flag.String("kuma", "", "uptime kuma url") // API base url
	interval  = flag.Int("interval", 60, "Interval between successive pings, in seconds")
)

func main() {
	flag.Parse()

	if *apiKey == "" || *targetUrl == "" || *apiUrl == "" {
		flag.Usage()
		return
	}

	for {
		startTime := time.Now()

		out, err := exec.Command("ping", "-c", "3", *targetUrl).Output()
		if err != nil {
			fmt.Printf("Error executing ping command: %v\n", err)
			return
		}

		pingDuration := time.Since(startTime)

		output := string(out)
		status := "down"
		if strings.Contains(output, "received") {
			status = "up"
		}

		packetLossRe := regexp.MustCompile(`(\d+)% packet loss`)
		packetLossMatch := packetLossRe.FindStringSubmatch(output)
		packetLoss := "unknown"
		if len(packetLossMatch) > 0 {
			packetLoss = packetLossMatch[1]
		}

		rttRe := regexp.MustCompile(`rtt min/avg/max/mdev = [\d\.]+/([\d\.]+)/[\d\.]+/[\d\.]+ ms`)
		rttMatch := rttRe.FindStringSubmatch(output)
		pingTime := "unknown"
		if len(rttMatch) > 0 {
			pingTime = rttMatch[1]
		}

		msg := fmt.Sprintf("Packet loss: %s%%", packetLoss)
		sendPingData(status, msg, pingTime)

		if *interval > int(pingDuration.Seconds()) {
			time.Sleep(time.Duration(*interval-int(pingDuration.Seconds())) * time.Second)
		}
	}
}

func sendPingData(status string, msg string, pingTime string) {
	fullUrl := fmt.Sprintf("%s/api/push/%s?status=%s&msg=%s&ping=%s", *apiUrl, *apiKey, url.QueryEscape(status), url.QueryEscape(msg), url.QueryEscape(pingTime))

	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("Error sending data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}

	fmt.Printf("Successfully sent status data: %s\n", status)
}
