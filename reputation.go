package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Reput struct {
	IpAddress            string   `json:"ipAddress"`
	IsPublic             bool     `json:"isPublic"`
	IpVersion            int      `json:"ipVersion"`
	IsWhitelisted        bool     `json:"isWhitelisted"`
	AbuseConfidenceScore int      `json:"abuseConfidenceScore"`
	CountryCode          string   `json:"countryCode"`
	UsageType            string   `json:"usageType"`
	Isp                  string   `json:"isp"`
	Domain               string   `json:"domain"`
	Hostnames            []string `json:"hostnames"`
	TotalReports         int      `json:"totalReports"`
	NumDistinctUsers     int      `json:"numDistinctUsers"`
	LastReportedAt       string   `json:"lastReportedAt"`
}

type Data struct {
	Reput Reput `json:"data"`
}

func GetReput(ip string, key string) (*Reput, error) {
	var data Data

	params := url.Values{}
	params.Add("ipAddress", ip)
	params.Add("maxAgeInDays", "90")

	req, err := http.NewRequest("GET", "https://api.abuseipdb.com/api/v2/check?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 || (!(strings.Contains(resp.Header.Get("Content-Type"), "application/json"))) {
		return nil, errors.New("Bad response from API : " + strconv.Itoa(resp.StatusCode) + "\nContent-Type : " + resp.Header.Get("Content-Type"))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &(data.Reput), nil
}
