package components

import (
	"io"
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

	return
}

func (me *Label) Update() {
	
}
	
func (me *Label) Size() (w, h int) {
	return
}
