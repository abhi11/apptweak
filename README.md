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
        auth := apptweak.NewAuth(token) // token to be generated from apptweak website
        params := apptweak.Parameters{Lang: "en", Country: "us", Type: "free"}
     	r := apptweak.NewTopAndroidAppSearchRequest(auth, params, "Business")
		resp, err := r.Run()
}
```
See [api_test.go](https://github.com/abhi11/apptweak/blob/master/api_test.go) for more details or check the wiki for documentation

To know more about apptweak, visit their [site](https://apptweak.io/) and read API Docs

#### TODO
* Add feature: Keywords Stats
* Write documentation
