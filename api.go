package apptweak

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"

	"net/http"
)

func httpRequest(method, url string, header http.Header) (io.ReadCloser, error) {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	r.Header = header
	c := &http.Client{Transport: tr} // skip certificate check
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error code: %d", resp.StatusCode))
	}

	return resp.Body, err

}

func httpGet(url string, header http.Header) (io.ReadCloser, error) {
	return httpRequest("GET", url, header)
}

// Wrapper for Getting TopApps from a category
// token is important to be passed
func GetTopAppsInCategory(category, country, lang, kind, token string) (AppResponse, error) {
	var appResp AppResponse
	auth := NewAuth(token)
	params := Parameters{Country: country, Lang: lang, Type: kind}
	req := NewTopAndroidAppSearchRequest(auth, params, category)
	appResp, err := req.Run()
	return appResp, err
}

func GetTopAppsForTerm(term, country, lang, token string) (AppResponse, error) {
	auth := NewAuth(token)
	params := Parameters{Term: term, Country: country, Lang: lang}
	req := NewAndroidAppTermSearchRequest(auth, params)

	return req.Run()
}
