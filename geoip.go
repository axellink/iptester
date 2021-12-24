package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type GeoIP struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
	Message     string  `json:"message"`
}

func Request(ip string) (*GeoIP, error) {
	var geoip GeoIP
	req, err := http.Get("http://ip-api.com/json/" + ip)
	// Check general error
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	// Check request status
	if req.StatusCode != 200 || (!(strings.Contains(req.Header.Get("Content-Type"), "application/json"))) {
		return nil, errors.New("Bad response from API : " + strconv.Itoa(req.StatusCode) + "\nContent-Type : " + req.Header.Get("Content-Type"))
	}
	// Get body into byte array
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	// Unmarshal JSON
	if !json.Valid(body) {
		return nil, errors.New("Unvalid JSON")
	}
	err = json.Unmarshal(body, &geoip)
	if err != nil {
		return nil, err
	}
	if geoip.Status != "success" {
		return nil, errors.New("Request failed : " + geoip.Message)
	}
	return &geoip, nil
}
