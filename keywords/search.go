package keywords

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const AMOUNT int = 100

func Search(email string) *[]string {
	url := fmt.Sprintf("https://www.google.com/search?num=%s&q=\"%s\"", strconv.Itoa(AMOUNT), email)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	var results []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}

				if !strings.HasPrefix(attr.Val, "/url?q=") {
					continue
				}

				attr.Val = strings.TrimPrefix(attr.Val, "/url?q=")
				if strings.HasPrefix(attr.Val, "https://accounts.google.com/") || strings.HasPrefix(attr.Val, "https://support.google.com/") {
					continue
				}
				results = append(results, attr.Val)
				break
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Return results, limiting to AMOUNT
	if len(results) > AMOUNT {
		results = results[:AMOUNT]
	}
	return &results
}
