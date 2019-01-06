package layout

import (
	"github.com/mattn/go-gtk/gtk"

	"../constant"
	"../helper"
)


func Start(){
	gtk.Init(nil)
	window := helper.CreateWindow(constant.WindowTitle,constant.WindowWidth,constant.WindowHeight)
	window.Add(CreateHeader())

	box := gtk.NewHBox(false,10)
	box.Add(CreateLeft())
	box.Add(CreateMain())
	window.Add(box)
	gtk.Main()
}



func CreateHeader() *gtk.Box{

	return nil
}


func CreateLeft() *gtk.Box{

	return nil
}

func CreateMain() *gtk.Box{

	return nil
}
