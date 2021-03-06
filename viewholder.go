package ext

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type (
	// ViewHolder ViewHolder
	ViewHolder struct {
		root  fyne.CanvasObject
		vlist map[string]fyne.CanvasObject
	}
)

func newViewHolder() *ViewHolder {
	return &ViewHolder{
		vlist: make(map[string]fyne.CanvasObject),
	}
}

// Add add a widget with a name
func (v *ViewHolder) Add(w fyne.CanvasObject, name string) {
	v.vlist[name] = w
}

// GetButton GetButton
func (v *ViewHolder) GetButton(name string) *widget.Button {
	obj, ok := v.vlist[name]
	if !ok {
		panic("Canvas Object with name [" + name + "] doesn't exist")
	}

	b, ok := obj.(*widget.Button)
	if !ok {
		panic("Canvas Object with name [" + name + "] is not button")
	}

	return b
}
