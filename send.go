package main

import (
	"fmt"
	"net/http"
	"net/url"
)

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
