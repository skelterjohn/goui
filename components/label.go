package components

import (
	"io"
	"fmt"
)

type Label struct {
	text string
}

func NewLabel(text string) (me *Label) {
	me = new(Label)
	me.text = text
	return
}

func (me *Label) SetText(text string) {
	
}

func (me *Label) Render(html io.Writer) (e error) {
	// write some html
	// connect click to the js on the other end, somehow
	_, e = fmt.Fprintf(html, "<label>\n")
	if e != nil { panic(e) }
	_, e = fmt.Fprintf(html, "%s\n", me.text)
	if e != nil { panic(e) }
	_, e = fmt.Fprintf(html, "</label>\n")
	if e != nil { panic(e) }

	return
}

func (me *Label) Update() {
	
}
	
func (me *Label) Size() (w, h int) {
	return
}
