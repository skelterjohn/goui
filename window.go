package goui

type Window struct {
	title string
	layout Layout
	w, h int
	dirty bool
	X <-chan struct{}
}

func NewWindow(title string) (me *Window) {
	me = new(Window)
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
}

func (me *Window) Center() {
	//only available if it's a private chrome instance that we can move around
}
