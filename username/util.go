package username

import (
	"net/http"
)

func checkStatusAndReturnUrl(url string, intentedStatus int) *string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
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
