package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GeolocationResponse struct {
	Result int `json:"result"`
	Data   struct {
		Lat   float64 `json:"lat"`
		Lng   float64 `json:"lon"`
		Range float64 `json:"range"`
		Time  int64   `json:"time"`
	} `json:"data"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Syntax: main.go <BSSID>")
		return
	}

	bssid := os.Args[1]

	// Send request to mylnikov.org WiFi geolocation API
	url := fmt.Sprintf("https://api.mylnikov.org/wifi?v=1.1&bssid=%s", bssid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse response
	var geolocationResponse GeolocationResponse
	err = json.NewDecoder(resp.Body).Decode(&geolocationResponse)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// Check result
	if geolocationResponse.Result != 200 {
		fmt.Println("Error: Result is not 200.")
		return
	}

	// Print data
	fmt.Printf("Latitude: %f\nLongitude: %f\nRange: %f\nTime: %d\n", geolocationResponse.Data.Lat, geolocationResponse.Data.Lng, geolocationResponse.Data.Range, geolocationResponse.Data.Time)
}
