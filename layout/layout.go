package layout

import (
	"github.com/mattn/go-gtk/gtk"

	"../common"
	"../helper"
	"../component"
)


func Start(){
	gtk.Init(nil)
	window := helper.CreateWindow(common.WindowTitle,common.WindowWidth,common.WindowHeight)

	vBox := gtk.NewVBox(false,0)
	vBox.SetSizeRequest(common.WindowWidth,common.WindowHeight)


	mBox := gtk.NewHBox(false,10)
	mBox.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	mBox.Add(CreateLeft())
	mBox.Add(CreateMain())

	vBox.Add(CreateHeader())
	vBox.Add(mBox)

	vBox.ShowAll()

	window.Add(vBox)
	gtk.Main()
}



func CreateHeader() *gtk.Fixed{

	header := gtk.NewFixed()
	header.SetSizeRequest(common.WindowWidth,common.NavItemHeight)

	hBox := gtk.NewHBox(true,5)

	connItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/conn.png","conn")
	docItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/doc.png","doc")
	indexItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/index.png","index")
	editItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/editDoc.png","edit")
	addItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/addDoc.png","add")
	searchItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/search.png","search")

	hBox.Add(connItem)
	hBox.Add(docItem)
	hBox.Add(indexItem)
	hBox.Add(editItem)
	hBox.Add(addItem)
	hBox.Add(searchItem)

	header.Put(hBox,20,10)
	header.ShowAll()
	return header
}


func CreateLeft() *gtk.VBox{
	left := gtk.NewVBox(false,0)
	left.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	return left
}

func CreateMain() *gtk.VBox{

	left := gtk.NewVBox(false,0)
	left.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.WindowLeftHeight)
	return left
}
