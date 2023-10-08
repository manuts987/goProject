package main

import (
	"fmt"
	"net"
	"github.com/manuts987/goProject/httpRequest"
	"strings"
)

func main() {
	var url string
	fmt.Println("enter the url(https://example.com or http://example.com) below:")
	fmt.Scanln(&url)
	hostname := getHostname(url)
	ip := getIp(hostname)
	fmt.Println(ip[1])

	// Teste si le module http fonctionne correctement
	bodyRq := httprequest.GetBodyRequest(url)
	headerRq := httprequest.GetheaderRequest(url)
	fmt.Println(bodyRq)
	fmt.Print(headerRq)
}

func getIp(url string) []net.IP {
	var ip []net.IP
	var err error
	ip, err = net.LookupIP(url)
	if err == nil {
		return ip
	}
	fmt.Println(err)
	return nil
}

func getHostname(url string) string {
	var hostname string = ""
	if strings.Contains(url, "http://") {
		hostname = strings.Trim(url, "http://")
	} else if strings.Contains(url, "https://") {
		hostname = strings.Trim(url, "https://")
	}else {
		return url
	}
	var newHostname string = ""
	if strings.Contains(url, "/") {
		newHostname = strings.Replace(hostname, "/", "", -1)
		return newHostname
	}
	return hostname
}
