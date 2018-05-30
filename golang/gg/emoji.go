package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
	goemoji "github.com/urakozz/go-emoji"
)

func main() {
	parseEmoji()
}

func emoji1() {
	fmt.Println("Hello World Emoji! 全身上下都是爱你的形状❤️ \U0001f44d")

	emoji.Println(":beer: Beer!!! 全身上下都是爱你的形状❤️ ")

	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)
}

func parseEmoji() {
	parser := goemoji.NewEmojiParser()
	var text = "a #💩 #and #🍦 #😳"
	replased := parser.ReplaceAllStringFunc(text, func(s string) string {
		return ""
	})

	fmt.Println(replased)
}
