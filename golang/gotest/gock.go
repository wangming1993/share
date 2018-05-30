package main

import (
	"net/http"

	"fmt"

	"gopkg.in/h2non/gock.v1"
)

func main() {

	defer gock.Off()

	gock.New("http://foo.com").
		Get("/bar").
		Reply(400).
		JSON(map[string]string{"foo": "bar"})

	res, err := http.Get("http://foo.com/bar")
	fmt.Println(res, err)
}
