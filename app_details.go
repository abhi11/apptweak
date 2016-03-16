package apptweak

import "encoding/json"

type Developer struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Website string `json:"website"`
}

type Ranking struct {
	CategoryId  string  `json:"category_id"`
	CountryCode string  `json:"country_code"`
	EndDate     string  `json:"end_date"`
	Ranks       []int64 `json:"ranks"`
	StartDate   string  `json:"start_date"`
	Type        string  `json:"type"`
}

type CountryRanking struct {
	CategoryId string           `json:"category_id"`
	Ranks      map[string]int64 `json:"ranks"`
	Type       string           `json:"type"`
}

type Rating struct {
	Average    float64          `json:"average"`
	Count      int64            `json:"count"`
	StartCount map[string]int64 `json:"start_count"`
}

func (r Rating) String() string {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type author struct {
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Profile string `json:"profile"`
}

type Review struct {
	Author  author `json:"author"`
	Body    string `json:"body"`
	Date    string `json:"date"`
	Id      string `json:"id"`
	Rating  int    `json:"rating"`
	Title   string `json:"title"`
	Version string `json:"version"`
}

func (rev Review) String() string {
	jsonBytes, err := json.Marshal(rev)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type Version struct {
	ReleaseDate  string `json:"release_date"`
	ReleaseNotes string `json:"release_notes"`
	Version      string `json:""version`
}

type StoreInformation struct {
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Genres           []string  `json:"genres"`
	Permissions      []string  `json:"permissions"`
	Price            string    `json:"price"`
	Screenshots      []string  `json:"screenshots"`
	Videos           []string  `json:"videos"`
	Slug             string    `json:"slug"`
	Title            string    `json:"title"`
	Versions         []Version `json:"versions"`
}

func (s StoreInformation) String() string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

type AppDetail struct {
	ApplicationId   string           `json:"application_id"`
	CountryCode     string           `json:"country_code"`
	Dev             Developer        `json:"developer"`
	Rankings        []Ranking        `json:"rankings"`
	CountryRankings []CountryRanking `json:"country_rankings"`
	Lang            string           `json:"language"`
	Ratings         Rating           `json:"ratings"`
	Reviews         []Review         `json:"reviews"`
	StoreInfo       StoreInformation `json:"store_info"`
}

func (a AppDetail) String() string {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// Find the Application Detail given an application id i.e package_name
type AppDetailRequest struct {
	ReqAuth Auth
	Params  Parameters
	Id      string
}

type AppDetailResponse struct {
	AD AppDetail `json:"content"`
}

func (r AppDetailRequest) Run() (AppDetailResponse, error) {
	url := "https://api.apptweak.com/android/applications/%s.json?"
	appResp := AppDetailResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}

	if r.Id == "" {
		return appResp, ApplicationIdNotPresent
	}

	// Make term, type empty as it is not required
	r.Params.Term = ""
	r.Params.Type = ""
	b, err := getRespForApplication(url, r.Id, r.ReqAuth.token, r.Params)
	if err != nil {
		return appResp, err
	}

	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	return appResp, nil
}

func NewAppDetailRequest(auth Auth, params Parameters, id string) AppDetailRequest {
	appReq := AppDetailRequest{ReqAuth: auth, Params: params, Id: id}
	return appReq
}

// This is exactly same as AppDetailsRequest
// Use it if only when ratings for an App is needed
type AppRatingRequest struct {
	ReqAuth Auth
	Params  Parameters
	Id      string
}

type AppRatingResponse struct {
	AppRating Rating `json:"content"`
}

func (r AppRatingRequest) Run() (AppRatingResponse, error) {
	url := "https://api.apptweak.com/android/applications/%s/ratings.json?"
	appResp := AppRatingResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}
	if r.Id == "" {
		return appResp, ApplicationIdNotPresent
	}

	// Make term, type empty as it is not required
	r.Params.Term = ""
	r.Params.Type = ""
	b, err := getRespForApplication(url, r.Id, r.ReqAuth.token, r.Params)
	if err != nil {
		return appResp, err
	}

	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	return appResp, nil
}

func NewAppRatingRequest(auth Auth, params Parameters, id string) AppRatingRequest {
	appReq := AppRatingRequest{ReqAuth: auth, Params: params, Id: id}
	return appReq
}

// This is exactly same as AppDetailsRequest
// Use it if only when store info for an App is needed
type AppStoreInfoRequest struct {
	ReqAuth Auth
	Params  Parameters
	Id      string
}

type AppStoreInfoResponse struct {
	StoreInfo StoreInformation `json:"content"`
}

func (r AppStoreInfoRequest) Run() (AppStoreInfoResponse, error) {
	url := "https://api.apptweak.com/android/applications/%s/information.json?"
	appResp := AppStoreInfoResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}
	if r.Id == "" {
		return appResp, ApplicationIdNotPresent
	}

	// Make term, type empty as it is not required
	r.Params.Term = ""
	r.Params.Type = ""
	b, err := getRespForApplication(url, r.Id, r.ReqAuth.token, r.Params)
	if err != nil {
		return appResp, err
	}

	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	return appResp, nil
}

func NewAppStoreInfoRequest(auth Auth, params Parameters, id string) AppStoreInfoRequest {
	appReq := AppStoreInfoRequest{ReqAuth: auth, Params: params, Id: id}
	return appReq
}

// This is exactly same as AppDetailsRequest
// Use it if only when reviews for an App is needed
type AppReviewsRequest struct {
	ReqAuth Auth
	Params  Parameters
	Id      string
}

type AppReviewResponse struct {
	Reviews []Review `json:"content"`
	Count   int      `json:"-"`
}

func (r AppReviewsRequest) Run() (AppReviewResponse, error) {
	url := "https://api.apptweak.com/android/applications/%s/reviews.json?"
	appResp := AppReviewResponse{}

	if r.ReqAuth.token == "" {
		return appResp, AuthNotPresent
	}
	if r.Id == "" {
		return appResp, ApplicationIdNotPresent
	}

	// Make term, type empty as it is not required
	r.Params.Term = ""
	r.Params.Type = ""
	b, err := getRespForApplication(url, r.Id, r.ReqAuth.token, r.Params)
	if err != nil {
		return appResp, err
	}

	err = bind(b, &appResp)
	if err != nil {
		return appResp, err
	}
	appResp.Count = len(appResp.Reviews)
	return appResp, nil
}

func NewAppReviewsRequest(auth Auth, params Parameters, id string) AppReviewsRequest {
	appReq := AppReviewsRequest{ReqAuth: auth, Params: params, Id: id}
	return appReq
}
