package main

import (
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/layouts"
	"github.com/skelterjohn/goui/components"
)

func main() {
	window := goui.NewWindow("Example")
	window.ReSize(500, 500)
	window.Center()
	
	dlayout := layouts.NewDialogLayout()
	window.Layout = dlayout
	
	label := components.NewLabel("Here is some text!")
	dlayout.SetMessage(label)
	okbutton := components.NewButton("Ok")
	dlayout.SetButton(okbutton)
	
	window.Show()

	<-okbutton.Click
	
	label.SetText("You just clicked")
	
	<-window.X
}
