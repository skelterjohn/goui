package components

import (
	"fmt"
	"io"
	"github.com/skelterjohn/goui/messages"
)

type Button struct {
	label string
	mname string
	Click <-chan []byte
}

func NewButton(label string) (me *Button) {
	me = new(Button)
	me.label = label
	me.mname, me.Click = messages.GetMessenger()

	return
}

func (me *Button) Render(html io.Writer) (e error) {
	// write some html
	// connect click to the js on the other end, somehow
	_, e = fmt.Fprintf(html, "<button label=\"%s\">\n", me.label)
	if e != nil { panic(e) }
	_, e = fmt.Fprintf(html, "<a href=/m/%s>click</a>\n", me.mname)
	_, e = fmt.Fprintf(html, "</button>\n")
	if e != nil { panic(e) }

	return
}
