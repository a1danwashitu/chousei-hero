package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, World!")
	registerCallbacks()
	<-c
}

func getCSV(hash string) string {
	url := "https://chouseisan.com/schedule/List/createCsv?h=" + hash

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := io.ReadAll(resp.Body)

	return string(byteArray)
}

func registerCallbacks() {
	js.Global().Set("output", js.FuncOf(outputChouseisan))
}

func outputChouseisan(this js.Value, args []js.Value) interface{} {
	hash := textToStr(args[0])
	text := getCSV(hash)
	fmt.Println("csv:", text)

	outputToHtml(text)
	return nil
}

func textToStr(v js.Value) string {
	return js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
}

func outputToHtml(text string) {
	println("print:", text)
	js.Global().Get("document").Call("getElementById", "output").Set("innerHTML", text)
}
