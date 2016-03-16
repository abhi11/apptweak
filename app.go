package apptweak

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type App struct {
	Developer string   `json:"developer"`
	Genres    []string `json:"genres"`
	Icon      string   `json:"icon"`
	Id        string   `json:"id"` // same as package_name
	Price     string   `json:"price"`
	Title     string   `json:"title"`
}

func (a App) String() string {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// Contains the response only if successful request,
// Conatins the count and the Apps
type AppResponse struct {
	Apps  []App `json:"content"`
	Count int   `json:"-"`
}

type Parameters struct {
	// Mandatory and only for TermSearch Call.
	// Do not use it for Top Search query
	Term string

	// Optional. Default:us
	Country string

	// Optional. Default:en
	Lang string

	// Optional. Default: free. Possible Values : free | paid | grossing
	// Use in TopSearchRequest
	// Do not use it in TermSearchRequest
	Type string
}

// To get the top 200 apps from a category
type TopAndroidAppSearchRequest struct {
	ReqAuth  Auth
	Params   Parameters
	Category string
}

func (r TopAndroidAppSearchRequest) Run() (AppResponse, error) {
	url := "https://api.apptweak.com/android/categories/%s/top.json?"
	appResp := AppResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}

	if r.Category == "" {
		return appResp, CategoryNotPresent
	}

	// Make term empty as it is not required
	r.Params.Term = ""
	fUrl := fmt.Sprintf(url, r.Category)
	fUrl = getUrlWithParams(fUrl, r.Params)
	header := make(http.Header)
	header.Add("X-Apptweak-Key", r.ReqAuth.token)

	b, err := httpGet(fUrl, header)
	if err != nil {
		return appResp, err
	}
	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	appResp.Count = len(appResp.Apps)
	return appResp, nil
}

// Returns a TopAppSerachRequest, call Run to get the result
// auth and category are mandatory
func NewTopAndroidAppSearchRequest(auth Auth, params Parameters, category string) TopAndroidAppSearchRequest {
	appReq := TopAndroidAppSearchRequest{}
	appReq.ReqAuth = auth
	appReq.Params = params
	appReq.Category = category
	return appReq
}

// Get the first 20 apps that would appear
// on google play search when the term is types
// here Term in parameters in manadatory
type AndroidAppTermSearchRequest struct {
	ReqAuth Auth
	Params  Parameters
}

func (r AndroidAppTermSearchRequest) Run() (AppResponse, error) {
	url := "https://api.apptweak.com/android/searches.json?"
	appResp := AppResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}

	// Make term empty as it is not required
	if r.Params.Term == "" {
		return appResp, TermNotPresent
	}

	url = getUrlWithParams(url, r.Params)
	header := make(http.Header)
	header.Add("X-Apptweak-Key", r.ReqAuth.token)

	b, err := httpGet(url, header)
	if err != nil {
		return appResp, err
	}
	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	appResp.Count = len(appResp.Apps)
	return appResp, nil
}

// Returns a AndroidAppTermSearchRequest, call Run to get the result
// auth and params.Term are mandatory
func NewAndroidAppTermSearchRequest(auth Auth, params Parameters) AndroidAppTermSearchRequest {
	appReq := AndroidAppTermSearchRequest{}
	appReq.ReqAuth = auth
	appReq.Params = params
	return appReq
}
