package goui

import (
	"fmt"
	"http"
	"io"
	"log"
	_ "github.com/skelterjohn/goui/server"
)

type Window struct {
	w, h int
	dirty bool

	path string
	serving bool

	title string
	layout Layout

	X <-chan struct{}
}

func NewWindow(title, path string) (me *Window) {
	me = new(Window)
	me.path = path
	me.SetTitle(title)
	me.SetSize(400, 300)
	me.dirty = true
	return
}

func (me *Window) SetTitle(title string) {
	me.title = title
	me.dirty = true
}

func (me *Window) SetLayout(layout Layout) {
	me.layout = layout
}

func (me *Window) SetSize(w, h int) {
	me.w, me.h = w, h
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
	_, e = fmt.Fprintf(html, "<window title=\"%s\">\n", me.title)
	if e != nil { panic(e) }
	if me.layout != nil {
		e = me.layout.Render(html)
		if e != nil { panic(e) }
	}
	_, e = fmt.Fprintf(html, "</window>\n")
	if e != nil { panic(e) }
	return
}

func (me *Window) Center() {
	//only available if it's a private chrome instance that we can move around
}
