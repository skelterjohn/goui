package layouts

import (
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/components"
)

type Dialog struct {
	message goui.Component
	button *components.Button
	dirty bool
}

func NewDialog() (me *Dialog) {
	me = new(Dialog)

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

func (me *Dialog) Update() {
	if me.dirty {
		// update the client-side
	}
	me.dirty = false
}