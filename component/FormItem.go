package component

import (
	"github.com/mattn/go-gtk/gtk"
)

const (
	FormItemWith = 160
	FormItemHeight = 30
)

func CreateFormItem(width int, height int,text string) *gtk.HBox{
	hBox := gtk.NewHBox(false,0)
	hBox.SetVisible(true)
	hBox.SetUSize(FormItemWith,FormItemHeight)

	label := gtk.NewLabel(text)
	label.SetUSize(FormItemWith/2,FormItemHeight)

	entry := gtk.NewEntry()
	entry.SetText(text)
	entry.SetUSize(FormItemWith/2,FormItemHeight)
	
	hBox.Add(label)
	hBox.Add(entry)

	return hBox
}
