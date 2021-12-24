package main

import (
	"log"
	"fmt"
	"os"
)

type Info struct {
	Geoip GeoIP
	Reput Reput
}

func GetColor(score int) (string) {
	switch {
	case score < 10 :
		return "\033[32m"
	case score < 80 :
		return "\033[33m"
	default:
		return "\033[31m"
	}
}

func PrintShort(info Info) {
	color := GetColor(info.Reput.AbuseConfidenceScore)
	fmt.Println(string(color),info.Reput.IpAddress,":",info.Geoip.Country,",",info.Geoip.City,"(",info.Reput.Domain,") ,Malicious :",info.Reput.AbuseConfidenceScore,string("\033[0m"))
}

func main() {
	config, err := GetConf()
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range os.Args[1:]{
		geoip, err := Request(ip)
		if err != nil {
			log.Fatal(err)
		}

		reput, err := GetReput(ip, config.Key)
		if err != nil {
			log.Fatal(err)
		}

		info := Info{Geoip: *geoip, Reput: *reput}
		PrintShort(info)
	}
}
