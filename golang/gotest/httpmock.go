package main

import (
	"net/http"

	"fmt"

	"github.com/jarcoal/httpmock"
)

func main() {
	fixture := `{"status":{"message": "Your message", "code": 200}}`
	responder, _ := httpmock.NewJsonResponder(200, fixture)
	fakeUrl := "https://api.mybiz.com/articles.json"
	httpmock.RegisterResponder("GET", fakeUrl, responder)

	resp, err := http.Get(fakeUrl)
	fmt.Println(resp, err)
}
