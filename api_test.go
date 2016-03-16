package apptweak

import (
	"fmt"
	"testing"
)

func TestGetTopAppsInCategory(t *testing.T) {
	token := "_ZSy8zFWOZuVhxLRuA9GEv_zJrs"
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
