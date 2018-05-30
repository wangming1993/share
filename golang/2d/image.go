package main

import (
	"fmt"
	"image"
	"net/http"
	"os"
	"time"

	"image/gif"
	"image/png"

	"image/jpeg"

	"io/ioutil"

	"bytes"

	"github.com/fogleman/gg"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

var index = 0

func main() {
	URL := "https://wx1.sinaimg.cn/large/0067A3aOgy1fr56gbpmv0g30780cthe1.gif"
	createImageFile(URL)
	URL = "https://mini-pics-10.bzy.ai/mirror/aHR0cHM6Ly93eDEuc2luYWltZy5jbi9sYXJnZS8wMDY1bFZXMGx5MWZybnJhMjhpcWpqMzBnbzBnb2duMy5qcGc=?bzytype=jpg&bzyw=600&bzyh=600"

	createImageFile(URL)
	URL = "https://mini-pics-1.bzy.ai/mirror/aHR0cHM6Ly93eDQuc2luYWltZy5jbi9sYXJnZS85MDU4YmUxZWd5MWZyM3FzeHpkOHlnMjBiNDA2OTRxcC5naWY=?bzytype=gif&bzyw=400&bzyh=225"
	createImageFile(URL)
	URL = "http://samples.fileformat.info/format/bmp/sample/dc59e50046b84768b5df4191ec16b9c3/LAND2.BMP?AWSAccessKeyId=0V91BEFA7GM093MEVMG2&Signature=aJV6OOyMBuI8m7f6TEA1za0SM3Q%3D&Expires=1527244686"
	createImageFile(URL)

	URL = "http://mini-pics.bzy.ai/mirror/aHR0cHM6Ly9pbWFnZXMyMDE1LmNuYmxvZ3MuY29tL25ld3NfdG9waWMvMjAxNjEyMTcwOTI5Mjk3MTQtMjUyNTUwOTkwLnBuZw=="
	createImageFile(URL)

	URL = "http://mini-pics.bzy.ai/mirror/aHR0cHM6Ly9tbWJpei5xcGljLmNuL21tYml6X2dpZi9MZDRmT1lLdTZ4YWhTRjNMOXRZd2lhMnRraDJRVm53eVJCSTZKSXpRYUE3ZVF1WjI2UUI5RGVwS01ueElyaWN0VnhFMWhWb2g1bnowU2hTUGVzZDh0clpBLzY0MD93eF9mbXQ9Z2lm"
	createImageFile(URL)

	URL = "http://mini-pics.bzy.ai/mirror/aHR0cHM6Ly91c2VyLWdvbGQtY2RuLnhpdHUuaW8vMjAxOC81LzI0LzE2MzkwNTA5YzgzYWVmMDM/aW1hZ2VWaWV3Mi8wL3cvMTI4MC9oLzk2MC9mb3JtYXQvd2VicC9pZ25vcmUtZXJyb3IvMQ=="
	createImageFile(URL)
}

func createImageFile(URL string) {
	index++
	img, err := downloadImage(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(fmt.Sprintf("img-%d.jpg", index))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = png.Encode(f, img)
	fmt.Println(err)
}

func draw() {
	start := time.Now()
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("out.png")

	end := time.Now()

	fmt.Println("cost time: ", end.Sub(start))
}

func downloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var img image.Image

	byteData, _ := ioutil.ReadAll(resp.Body)

	reader := bytes.NewReader(byteData)
	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		return nil, err
	}

	reader = bytes.NewReader(byteData)

	switch format {
	case "png":
		img, err = png.Decode(reader)
	case "jpeg":
		img, err = jpeg.Decode(reader)
	case "gif":
		img, err = gif.Decode(reader)
	case "bmp":
		img, err = bmp.Decode(reader)
	case "webp":
		img, err = webp.Decode(reader)
	case "tiff":
		img, err = tiff.Decode(reader)
	}

	return img, err
}

// Guess image format from gif/jpeg/png/webp
func guessImageFormat(url string) (format string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	_, format, err = image.DecodeConfig(resp.Body)
	fmt.Println(format, err)
	return
}
