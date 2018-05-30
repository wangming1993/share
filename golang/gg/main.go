package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	_ "net/http/pprof"

	"github.com/bzy-ai/rpc-mesh/proto"
	"github.com/bzy-ai/rpc-mesh/sdk/gg"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// netstat -a -p TCP -n | grep 9090

var (
	host = flag.String("host", "127.0.0.1", "listening host")
	port = flag.Int("port", 8080, "listening port")
)

func main() {
	flag.Parse()

	router := gin.Default()

	router.GET("/", drawFeed)

	go func() {
		http.ListenAndServe(":6062", nil)
	}()

	go trigger()

	router.Run(fmt.Sprintf("%s:%d", *host, *port))
}

func trigger() {
	time.Sleep(10 * time.Second)

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			go func() {
				for i := 0; i < rand.Intn(5); i++ {
					resp, err := http.Get("http://127.0.0.1:8080/")
					fmt.Println(resp, err)
				}
			}()
		}
	}
}

func PrintMemUsage() {
	return
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func drawFeed(c *gin.Context) {

	fmt.Println("before:")
	PrintMemUsage()

	miniAppURL := fmt.Sprintf("https://api.bzy.ai/circles/qrcode/%s-%s-%s", "dev", "4", "91061")

	feed := proto.Feed{
		Id:   int64(91061),
		Type: "image",
		//Text: `很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字很多很多的字`,
		Text: "全身上下都是爱你的形状    ❤️ ​",
		Images: []string{
			"http://pic16.nipic.com/20110909/2238196_094518585165_2.jpg",
			//"http://pic16.nipic.com/20110909/2238196_094518585165_2.jpg",
			"https://wx2.sinaimg.cn/large/9c01a767ly1fqk6pfybxoj21dc0ww4h7.jpg",
			//"https://wx1.sinaimg.cn/large/9c01a767ly1fqk6pgf7vnj21dc0wwat9.jpg",
			//"https://wx2.sinaimg.cn/large/9c01a767ly1fqk6pfemdvj20ww1dc4hl.jpg",
			"https://wx3.sinaimg.cn/large/9c01a767ly1fqk6ph0acij20ww1dch76.jpg",
		},
		Partner: &proto.FeedPartner{
			Id:          int64(1099),
			Name:        "三文鱼家的鱼",
			Avatar:      "https://cdn.bzy.ai/avatar/5ad586d2b558f.jpg",
			Description: "看些大海 也写些字",
		},
	}

	// rxEmoji := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	// feed.Text = rxEmoji.ReplaceAllString(feed.Text, "")
	// fmt.Println(feed.Text)

	image, err := gg.DrawTop100Feed(c, miniAppURL, feed)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		c.Writer.WriteHeader(400)

		return
	}
	c.Header("Content-Type", "image/jpg")
	c.Writer.Write(image)

	fmt.Println("after")
	PrintMemUsage()

	runtime.GC()

	fmt.Println("GC")
	PrintMemUsage()
}

func drawImage(imageURL string, width, height int) error {
	if !strings.Contains(imageURL, "bzy.ai") {
		imageURL = strings.Replace(imageURL, "://", "/", -1)
		imageURL = "https://mini-pics.bzy.ai/cache/" + imageURL
	}

	img, err := downloadImage(imageURL)
	if err != nil {
		return err
	}
	saveJPEG(img, "original.jpg")
	fmt.Println(img.Bounds().Dx(), img.Bounds().Dy())

	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()

	rx := float32(dx) / float32(width)
	ry := float32(dy) / float32(height)
	fmt.Println(rx, ry)
	if rx > ry {
		rx = ry
	}

	rWidth := int(float32(dx) / rx)
	rHeight := int(float32(dy) / rx)
	fmt.Println(rWidth, rHeight)
	img = imaging.Resize(img, rWidth, rHeight, imaging.Lanczos)
	saveJPEG(img, "resize.jpg")
	img = imaging.CropCenter(img, width, height)
	saveJPEG(img, "crop.jpg")

	fmt.Println(width, height)

	return nil
}

func downloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	avatar, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return avatar, nil
}

func saveJPEG(m image.Image, name string) {
	out, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
