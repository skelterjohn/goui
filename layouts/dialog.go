package layouts

import (
	"template"
	"io"
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/components"
)

const DialogHTMLTemplateFormat = `
Dialog
<br>
{{ComponentWriter .Message}}
<br>
{{ComponentWriter .Button}}
`

var DialogHTMLTemplate = template.Must(goui.ParseExecTemplate(DialogHTMLTemplateFormat))

type DialogData struct {
	Message       goui.Component
	Button        *components.Button
	Width, Height int
}

type Dialog struct {
	data DialogData

	dirty bool
}

func NewDialog() (me *Dialog) {
	me = new(Dialog)

	return
}

func (me *Dialog) SetData(data DialogData) {
	me.data = data
	me.dirty = true
}

func (me *Dialog) GetData() DialogData {
	return me.data
}

func (me *Dialog) SetSize(w, h int) {
	me.data.Width, me.data.Height = w, h
	me.dirty = true
}

func (me *Dialog) Size() (w, h int) {
	w, h = me.data.Width, me.data.Height
	return
}

func (me *Dialog) Render(html io.Writer) (e error) {
	e = DialogHTMLTemplate.Execute(html, me.data)
	return
}

func (me *Dialog) Update() {
	if me.dirty {
		// update the client-side
	}
	me.dirty = false
}
