package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, World!")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("output", js.FuncOf(output))
}

func output(this js.Value, args []js.Value) interface{} {
	text := textToStr(args[0])

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
