package main

import (
	"time"

	"honnef.co/go/js/dom"
)

func main() {
	dom.GetWindow().AddEventListener("DOMContentLoaded", false, func(e dom.Event) {
		started := time.Now()
		screen := dom.GetWindow().Document().GetElementByID("screen").(*dom.HTMLCanvasElement)

		context := screen.GetContext2d() //Call("getContext", "2d")
		context.Font = "38pt Arial"
		context.FillStyle = "cornflowerblue"
		context.StrokeStyle = "blue"

		context.Call("fillText", "Hello Canvas", screen.Width/2-150, screen.Height/2-15)
		context.Call("strokeText", "Hello Canvas", screen.Width/2-150, screen.Height/2-15)
		println("total time taken is:", time.Since(started).String())
	})

}
