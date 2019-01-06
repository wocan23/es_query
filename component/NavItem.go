package component

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdkpixbuf"
)


/**
	创建一个带图标文字带按钮
 */
func GenerateNavItem(width int,height int,imagePath string,title string) *gtk.Button{
	button := gtk.NewButton()
	button.SetSizeRequest(width,height)

	vBox := gtk.NewVBox(false,0)
	vBox.SetSizeRequest(width,height)

	srcpixBuf,_ := gdkpixbuf.NewPixbufFromFile(imagePath)
	pixBuf := srcpixBuf.ScaleSimple(30,30,gdkpixbuf.INTERP_HYPER)
	image := gtk.NewImageFromPixbuf(pixBuf)
	image.SetSizeRequest(width,width)

	label := gtk.NewLabel(title)
	label.SetSizeRequest(width,height-width)
	label.Show()

	vBox.Add(image)
	vBox.Add(label)
	vBox.Show()

	button.Add(vBox)
	button.ShowAll()
	return button
}