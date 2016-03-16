package apptweak

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}

// Expects an url and append it with the Paramaters
func getUrlWithParams(url string, params Parameters) string {
	p := map[string]string{}
	if params.Term != "" {
		p["term"] = params.Term
	}
	if params.Country != "" {
		p["country"] = params.Country
	}
	if params.Lang != "" {
		p["language"] = params.Lang
	}
	if params.Type != "" {
		p["type"] = params.Type
	}
	plist := []string{}
	for k, v := range p {
		plist = append(plist, fmt.Sprintf("%s=%s", k, v))
	}
	if len(plist) < 1 {
		return url
	}
	url = url + strings.Join(plist, "&")
	return url
}

func getRespForApplication(url, id, token string, p Parameters) (io.ReadCloser, error) {
	fUrl := fmt.Sprintf(url, id)
	fUrl = getUrlWithParams(fUrl, p)
	header := make(http.Header)
	header.Add("X-Apptweak-Key", token)
	return httpGet(fUrl, header)
}
