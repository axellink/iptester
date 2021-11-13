package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GeoIP struct {
	status      string  `json:"status"`
	country     string  `json:"country"`
	countryCode string  `json:"countryCode"`
	region      string  `json:"region"`
	regionName  string  `json:"regionName"`
	city        string  `json:"city"`
	zip         string  `json:"zip"`
	lat         float64 `json:"lat"`
	lon         float64 `json:"lon"`
	timezone    string  `json:"timezone"`
	isp         string  `json:"isp"`
	org         string  `json:"org"`
	as          string  `json:"as"`
	query       string  `json:"query"`
}

func Request(ip string) (GeoIP, error) {
	var geoip GeoIP
	req, err := http.Get("http://ip-api.com/json/" + ip)
	// Check general error
	if err != nil {
		return geoip, err
	}
	defer req.Body.Close()
	// Check request status
	if req.StatusCode != 200 || (!(strings.Contains(req.Header.Get("Content-Type"), "application/json"))) {
		return geoip, errors.New("Bad response from API")
	}
	// Get body into byte array
	body, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(body[:]))
	if err != nil {
		return geoip, err
	}
	// Unmarshal JSON
	if !json.Valid(body) {
		return geoip, errors.New("Unvalid JSON")
	}
	err = json.Unmarshal(body, &geoip)
	if err != nil {
		return geoip, err
	}
	fmt.Println(geoip)
	return geoip, nil
}

func test() {
	geoip, err := Request("4.4.4.4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(geoip.status)
	fmt.Println(geoip.country)
	fmt.Println(geoip.query)
}

func main() {
	test()
}
