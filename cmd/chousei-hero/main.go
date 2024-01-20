package main

import (
	"fmt"
	"syscall/js"

	"github.com/a1danwashitu/chousei-hero/io"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, World!")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("output", js.FuncOf(outputChouseisan))
}

func outputChouseisan(this js.Value, args []js.Value) interface{} {
	text := textToStr(args[0])

	m, d, c := io.ReadChouseisan(text)

	outputToHtml("duties", d)
	outputToHtml("members", m)
	outputToHtml("status", c)

	return nil
}

func textToStr(v js.Value) string {
	return js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
}

func outputToHtml(id, text string) {
	println("print:", text)
	js.Global().Get("document").Call("getElementById", "output").Set("innerHTML", text)
}
