package username

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func checkStatusAndReturnUrl(url string, intentedStatus int) *string {
	resp, err := http.Get(url)
	//client := &http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		return http.ErrUseLastResponse
	//	},
	//}
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v\n", err)
		}
	}(resp.Body)
	if resp.StatusCode != intentedStatus {
		return nil
	}
	return &url
}

func csAndReturnCustomURL(url string, intendedStatus int, returnUrl *string) *string {
	resp := checkStatusAndReturnUrl(url, intendedStatus)
	if resp != nil {
		return returnUrl
	}
	return nil
}

func checkTextAndReturnUrl(url string, text string) *string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v\n", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	if strings.Contains(string(body), text) {
		return &url
	}
	return nil
}

func ctAndReturnCustomURL(url string, text string, returnUrl *string) *string {
	resp := checkTextAndReturnUrl(url, text)
	if resp != nil {
		return returnUrl
	}
	return nil
}
