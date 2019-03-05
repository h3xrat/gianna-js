package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"./handler"
	"./utils"
	"github.com/fatih/color"
)

func verifyDomain(domain string) bool {
	if domain == "" {

		return false
	}
	return true
}

func appendHTTP(domain string) string {
	if !strings.Contains(domain, "http://") {
		domain := "http://" + domain
		return domain
	}
	return domain
}

func processFile(wordlist string) {

	var domain string

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain = scanner.Text()
		if verifyDomain(domain) {
			domain = appendHTTP(domain)
		}
		g := handler.HTTPHandler{domain}
		g.SendRequest()
	}
}

func main() {

	var (
		domain   string
		wordlist string
	)

	color.Green(utils.Banner())
	flag.StringVar(&domain, "domain", "", "[+] Name of the target [+]")
	flag.StringVar(&wordlist, "wordlist", "", "[+] List of domains")
	flag.Parse()

	if verifyDomain(domain) {
		domain = appendHTTP(domain)
	}

	if domain != "" {
		gi := handler.HTTPHandler{domain}
		gi.SendRequest()
	}
	if wordlist != "" {
		processFile(wordlist)
	}

}
