package layouts

import (
	"template"
	"io"
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/components"
)

const DialogHTMLTemplateFormat = `
Dialog
{{ComponentWriter .Message}}
{{ComponentWriter .Button}}
`

var DialogHTMLTemplate = template.Must(goui.ParseExecTemplate(DialogHTMLTemplateFormat))

type DialogData struct {
	Message goui.Component
	Button *components.Button
}

type Dialog struct {
	dd DialogData

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
	me.dd.Message = message
	me.dirty = true
}

func (me *Dialog) SetButton(button *components.Button) {
	me.dd.Button = button
	me.dirty = true
}

func (me *Dialog) Render(html io.Writer) (e error) {
	e = DialogHTMLTemplate.Execute(html, me.dd)
	return
}

func (me *Dialog) Update() {
	if me.dirty {
		// update the client-side
	}
	me.dirty = false
}