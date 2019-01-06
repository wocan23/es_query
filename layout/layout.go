package layout

import (
	"github.com/mattn/go-gtk/gtk"

	"../constant"
	"../helper"
	"../component"
)


func Start(){
	gtk.Init(nil)
	window := helper.CreateWindow(constant.WindowTitle,constant.WindowWidth,constant.WindowHeight)

	vBox := gtk.NewVBox(false,10)

	mBox := gtk.NewHBox(false,10)
	mBox.Add(CreateLeft())
	mBox.Add(CreateMain())

	vBox.Add(CreateHeader())
	vBox.Add(mBox)
	
	window.Add(vBox)
	gtk.Main()
}



func CreateHeader() *gtk.HBox{

	header := gtk.NewHBox(true,10)

	connItem := component.GenerateNavItem(constant.NavItemWidth,constant.WindowHeight,"../images/conn.png","conn")
	queryItem := component.GenerateNavItem(constant.NavItemWidth,constant.WindowHeight,"../images/search.png","search")
	header.Add(connItem)
	header.Add(queryItem)

	return header
}


func CreateLeft() *gtk.Box{

	return nil
}

func CreateMain() *gtk.Box{

	return nil
}
