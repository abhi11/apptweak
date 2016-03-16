package apptweak

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func httpGet(url string, header http.Header) (io.ReadCloser, error) {
	r, err := http.NewRequest("GET", url, nil)
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

func GetTopAppsInCategory(category, country, lang, kind, token string) (AppResponse, error) {
	url := "https://api.apptweak.com/android/categories/%s/top.json?country=%s&language=%s&type=%s"
	fUrl := fmt.Sprintf(url, category, country, lang, kind)
	header := make(http.Header)
	header.Add("X-Apptweak-Key", token)
	appResp := AppResponse{}
	b, err := httpGet(fUrl, header)
	if err != nil {
		return appResp, err
	}
	err = Bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	appResp.Count = len(appResp.Apps)
	return appResp, nil
}
