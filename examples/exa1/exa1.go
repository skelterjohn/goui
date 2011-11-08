package main

import (
	"fmt"
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/layouts"
	"github.com/skelterjohn/goui/components"
)

func main() {

	label := components.NewLabel("Here is some text!")
	okbutton := components.NewButton("Ok")

	dlayout := layouts.NewDialog()
	dlayout.SetData(
		layouts.DialogData{
			Message: label,
			Button:  okbutton,
		})

	window := goui.NewWindow("Example 1", "exa1")
	window.SetData(
		goui.WindowData{
			Width:  500,
			Height: 500,
			Layout: dlayout,
		})

	window.Show()

	fmt.Println(<-okbutton.Click)

	<-window.X
}
