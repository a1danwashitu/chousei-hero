package main

import (
	"syscall/js"

	"github.com/a1danwashitu/chousei-hero/io"
	"github.com/a1danwashitu/chousei-hero/solve"
)

func main() {
	c := make(chan struct{})

	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("output", js.FuncOf(outputChouseisan))
	js.Global().Set("output2", js.FuncOf(exe))
}

func outputChouseisan(this js.Value, args []js.Value) interface{} {
	text := textToStr(args[0])

	m, d, c := io.ReadChouseisan(text)

	outputToHtml("duties", d)
	outputToHtml("members", m)
	outputToHtml("status", c)

	return nil
}

func exe(this js.Value, args []js.Value) interface{} {
	duties := textToStr(args[0])
	members := textToStr(args[1])
	statuses := textToStr(args[2])

	eventConf := io.ReadConfStrings(members, duties, statuses)

	assigns := solve.Solve(eventConf)

	assignsText := io.OutputResult(assigns)

	outputToHtml("result", assignsText)

	return nil
}

func textToStr(v js.Value) string {
	return js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
}

func outputToHtml(id, text string) {
	js.Global().Get("document").Call("getElementById", id).Set("innerHTML", text)
}
