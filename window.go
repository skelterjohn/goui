package goui

import (
	"fmt"
	"http"
	"template"
	"io"
	"log"
	_ "github.com/skelterjohn/goui/server"
)

const WindowHTMLTemplateFormat = `
<html>
<title>{{.Title}}</title>
<body>
I wish I was {{.Width}}x{{.Height}}
<br>
{{ComponentWriter .Layout}}
</body>
</html>
`

var WindowHTMLTemplate = template.Must(ParseExecTemplate(WindowHTMLTemplateFormat))

type WindowData struct {
	Title         string
	Layout        Layout
	Width, Height int
	Color         string
}

type Window struct {
	dirty bool

	path    string
	serving bool

	data WindowData

	X <-chan struct{}
}

func NewWindow(title, path string) (me *Window) {
	me = new(Window)
	me.path = path
	me.data.Title = title
	me.SetSize(400, 300)
	me.dirty = true
	return
}

func (me *Window) SetData(data WindowData) {
	me.data = data
	me.dirty = true
}

func (me *Window) GetData() WindowData {
	return me.data
}

func (me *Window) SetSize(w, h int) {
	me.data.Width, me.data.Height = w, h
	me.dirty = true
}

func (me *Window) Update() {
	if me.dirty {
		//send a message to the client
	}
	me.dirty = false
}

func (me *Window) Show() {
	//start serving
	//open the window
	if !me.serving {
		me.serve()
		me.serving = true
	}
}

func (me *Window) serve() {
	h := func(rw http.ResponseWriter, rq *http.Request) {
		log.Printf("window %s got request\n", me.path)
		e := me.Render(rw)
		if e != nil {
			panic(e)
		}
	}

	pattern := fmt.Sprintf("/w/%s", me.path)
	http.DefaultServeMux.HandleFunc(pattern, h)
	log.Println("Window listening at", pattern)
}

func (me *Window) Render(html io.Writer) (e error) {
	e = WindowHTMLTemplate.Execute(html, me.data)
	return
}

func (me *Window) Center() {
	//only available if it's a private chrome instance that we can move around
}
