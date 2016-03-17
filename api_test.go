package apptweak

import (
	"fmt"
	"os"
	"testing"
)

func TestGetTopAppsInCategory(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	resp, err := GetTopAppsInCategory("Lifestyle", "us", "en", "free", token)
	if err != nil {
		fmt.Println("error getting apps: ", err)
		t.Fail()
	}
	fmt.Println("Count: ", resp.Count)
	for _, app := range resp.Apps {
		fmt.Println(app.String())
	}
}

func TestGetTopAndroidAppsInCategory(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewTopAndroidAppSearchRequest(auth, Parameters{Lang: "en", Country: "us", Type: "free"}, "Business")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("error getting apps: ", err)
		t.Fail()
	}
	fmt.Println("Count: ", resp.Count)
	for _, app := range resp.Apps {
		fmt.Println(app.String())
	}
}

func TestGetAndroidTermSearch(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAndroidAppTermSearchRequest(auth, Parameters{Term: "happy"})
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAndroidTermSearch::error getting apps: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAndroidTermSearch::Count: ", resp.Count)
	for i, app := range resp.Apps {
		fmt.Println("TestGetAndroidTermSearch::", i, app.String())
	}
}

func TestGetAppDetails(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppDetailRequest(auth, Parameters{}, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppDetails::error getting app details: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppDetails::", resp.AD.String())
}

func TestGetAppRating(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppRatingRequest(auth, Parameters{}, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppRating::error getting app rating: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppRating::", resp.AppRating.String())
}

func TestGetAppReviews(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppReviewsRequest(auth, Parameters{}, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppReviews::error getting app rating: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppReviews:: count", resp.Count)
	for i, r := range resp.Reviews {
		fmt.Println("TestGetAppReviews::", i, r.String())
	}
}

func TestGetAppStoreInfo(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppStoreInfoRequest(auth, Parameters{}, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppStoreInfo::error getting app rating: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppStoreInfo::", resp.StoreInfo.String())
}

func TestGetAppTrends(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppTrendsRequest(auth, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppTrends::error getting app rating: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppTrends:: count", resp.Count)
	for i, r := range resp.AppTrends {
		fmt.Println("TestGetAppTrends::", i, r.String())
	}
}

func TestGetAppKeywordsRank(t *testing.T) {
	token := os.Getenv("APPTWEAK_TOKEN")
	auth := NewAuth(token)
	r := NewAppKeywordsRankRequest(auth, Parameters{}, "com.facebook.katana")
	resp, err := r.Run()
	if err != nil {
		fmt.Println("TestGetAppKeywordsRank::error getting app rating: ", err)
		t.Fail()
	}
	fmt.Println("TestGetAppKeywordsRank:: count", resp.Count)
	for i, r := range resp.Rankings {
		fmt.Println("TestGetAppKeywordsRank::", i, r.String())
	}
}
