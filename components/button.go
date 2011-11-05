package components

import (
	"io"
)

type Button struct {
	label string
	click chan struct{}
	Click <-chan struct{}
}

func NewButton(label string) (me *Button) {
	me = new(Button)
	me.label = label
	me.click = make(chan struct{})
	me.Click = me.click

	return
}

func (me *Button) Render(html io.Writer) (e error) {
	// write some html
	// connect click to the js on the other end, somehow

	return
}
