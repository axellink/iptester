# iptester
A little tool to test IP addresses quickly against a geolocation and a reputation API (respectively [ip-api](https://ip-api.com/) and [abuseIPDB](https://www.abuseipdb.com/))

## Dependency
I only use native Go libraries so no dependency for build.

APIs on the other end are different, ip-api doesn't need a key for its free access, but be aware that it's HTTP only (no SSL). AbuseIPDB needs a key, you can get it by creating a free account on the website.

## Build
Just build it the old Go way :
```
$ go build -o iptester
```

## Configure
ip-api doesn't need a key for free access so there is no configuration needed for it, but be aware the it will use HTTP (no SSL) for request in its free access plan.

AbuseIPDB on the other hand, needs a key, you can get one by registering a free account on their website. Once you have it, copy `key.json.tpl` to `key.json` and put your key in it.

It should look like this :
```
{
    "key":"YOUR-ABUSEIPDB-KEY"
}
```

## Usage
```
Usage   : ./iptester [options] -f ip_file|ip...
Options : -v : Verbose mode
          -h : Print this help and quits
```

Exemples :
* To check IPs from a file named `ips` (one IP per line) : `./iptester -f ips`
* To check IPs in verbose mode : `./iptester -v 8.8.8.8 8.8.4.4 1.1.1.1`
