package layouts

import (
	"io"
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/components"
)

type Dialog struct {
	message goui.Component
	button *components.Button
	dirty bool
	w, h int
}

func NewDialog() (me *Dialog) {
	me = new(Dialog)

	return
}

func (me *Dialog) SetSize(w, h int) {
	me.w, me.h = w, h
}

func (me *Dialog) Size() (w, h int) {
	w, h = me.w, me.h
	return
}

func (me *Dialog) SetMessage(message goui.Component) {
	me.message = message
	me.dirty = true
}

func (me *Dialog) SetButton(button *components.Button) {
	me.button = button
	me.dirty = true
}

func (me *Dialog) Render(html io.Writer) (e error) {
	return
}

func (me *Dialog) Update() {
	if me.dirty {
		// update the client-side
	}
	me.dirty = false
}