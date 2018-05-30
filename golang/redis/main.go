package main

import (
	"fmt"

	"github.com/wangming1993/share/golang/redis/client"
)

func main() {
	//client.Redis.Set("name", "mike")
	v, err := client.Redis.Get("name")
	fmt.Println(v, err)
}
