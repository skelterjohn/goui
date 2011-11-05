package goui

type Layout interface {
	Component
	//called by the container 
	SetSize(w, h int)

}
