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
