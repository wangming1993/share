package main

//////////////////////////////////////////////////////
// brew install phantomjs
// go get -u github.com/benbjohnson/phantomjs
/////////////////////////////////////////////////////

import (
	"fmt"
	"os"
	"time"

	"github.com/benbjohnson/phantomjs"
)

func main() {
	start := time.Now()
	err := workAsPage()
	fmt.Println(err)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func workAsPage() error {
	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()

	page, err := phantomjs.DefaultProcess.CreateWebPage()
	if err != nil {
		return err
	}
	defer page.Close()

	// Open a URL.
	if err := page.Open("./test.html"); err != nil {
		return err
	}

	// Setup the viewport and render the results view.
	if err := page.SetViewportSize(512, 500); err != nil {
		return err
	}

	return page.Render("html.png", "png", 100)
}
