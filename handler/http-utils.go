package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//HTTPHandler struct to handler http requests
type HTTPHandler struct {
	URL string
}

//SendRequest sending request to a domain
func (g *HTTPHandler) SendRequest() {
	domain := g.URL
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	request, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		log.Println("[+] Error Sending request to: ", domain)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	//make request
	response, err := client.Do(request)
	if err != nil {
		log.Println("[+] Erro get response from request to: ", domain)
	}
	defer response.Body.Close()

	ParserDomain(response)
}

//ParserDomain get all js files from a give domain
func ParserDomain(bodyDomain *http.Response) {
	doc, err := goquery.NewDocumentFromReader(bodyDomain.Body)
	if err != nil {
		fmt.Println("[+] Erro get information from ", bodyDomain.Request)
	}
	doc.Find("script").Each(func(index int, s *goquery.Selection) {
		js, _ := s.Attr("src")
		if js != "" {
			if strings.HasPrefix(js, "http://") || strings.HasPrefix(js, "https://") || strings.HasPrefix(js, "//") {
				fmt.Println(js)
			}
		}
	})
}
