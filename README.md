### Golang SDK for [apptweak.io](https://apptweak.io/)

#### Getting apptweak
```
go get -u github.com/abhi11/apptweak
```

#### Using apptweak
```go
import "github.com/abhi11/apptweak"

func main() {
	    token := os.Getenv("APPTWEAK_TOKEN")
        auth := NewAuth(token) // token to be generated from apptweak website
        params := Parameters{Lang: "en", Country: "us", Type: "free"}
     	r := apptweak.NewTopAndroidAppSearchRequest(auth, , "Business")
		resp, err := r.Run()
}
```
See api_test.go for more details or check the wiki for documentation

#### To know more about apptweak, visit their [site](https://apptweak.io/) and read API Docs

#### TODO
* Add feature: Wrapper for Application, App Keywords, App Trends
* Write documentation
