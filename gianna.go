package main

import (
	"flag"
	"strings"

	"./handler"
	"./utils"
	"github.com/fatih/color"
)

func verifyDomain(domain string) bool {
	if domain == "" {
		color.Green(utils.Banner())
		return false
	}
	return true
}

func appendHTTPS(domain string) string {
	color.Green(utils.Banner())
	if !strings.Contains(domain, "https://") || !strings.Contains(domain, "http://") {
		color.Red(" [+] U dont passed any http or https schema.. all requests will be made with https [+] ")
		domain := "https://" + domain
		return domain
	}
	return domain
}

func main() {

	var domain string
	flag.StringVar(&domain, "domain", "", "[+] Name of the target [+]")
	flag.Parse()

	if verifyDomain(domain) {
		domain = appendHTTPS(domain)
	}
	color.Yellow(domain)
	gi := handler.HTTPHandler{domain}
	gi.SendRequest()
}
