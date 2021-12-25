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
	fmt.Println(color,info.Reput.IpAddress,":",info.Geoip.Country,",",info.Geoip.City,"(",info.Reput.Domain,") ,Malicious :",info.Reput.AbuseConfidenceScore,"\033[0m")
}

func PrintLong(info Info) {
	color := GetColor(info.Reput.AbuseConfidenceScore)
	fmt.Println(color, "IP Address :", info.Reput.IpAddress,"\033[0m")
	fmt.Println(color, "Country :", info.Geoip.Country,"\033[0m")
	fmt.Println(color, "Country Code :", info.Geoip.CountryCode,"\033[0m")
	fmt.Println(color, "Region :", info.Geoip.Region,"\033[0m")
	fmt.Println(color, "City :", info.Geoip.City,"\033[0m")
	fmt.Println(color, "ISP :", info.Geoip.Isp,"\033[0m")
	fmt.Println(color, "Org :", info.Geoip.Org,"\033[0m")
	fmt.Println(color, "Score :", info.Reput.AbuseConfidenceScore,"\033[0m")
	fmt.Println(color, "Whitelisted :", info.Reput.IsWhitelisted,"\033[0m")
	fmt.Println(color, "Reports :", info.Reput.TotalReports,"\033[0m")
	fmt.Println(color, "Domain :", info.Reput.Domain,"\033[0m")
	fmt.Println(color, "Hostnames :","\033[0m")
	for _, hn := range(info.Reput.Hostnames) {
		fmt.Println(color, "  -", hn,"\033[0m")
	}
	fmt.Println()
}

func main() {
	config, err := GetConf()

	Print := PrintShort
	argsBegin := 1

	if os.Args[1] == "-v" {
		Print = PrintLong
		argsBegin = 2
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range os.Args[argsBegin:]{
		geoip, err := Request(ip)
		if err != nil {
			log.Fatal(err)
		}

		reput, err := GetReput(ip, config.Key)
		if err != nil {
			log.Fatal(err)
		}

		info := Info{Geoip: *geoip, Reput: *reput}
		Print(info)
	}
}
