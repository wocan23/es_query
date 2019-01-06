package helper

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func CreateWindow(title string,width int,height int) *gtk.Window{
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle(title)
	window.SetIconName(title)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "foo")

	window.SetSizeRequest(width, height)
	window.ShowAll()

	return window
}

func CreateImage(width int,height int,imagePath string) *gtk.Image{
	srcpixBuf,_ := gdkpixbuf.NewPixbufFromFile(imagePath)
	pixBuf := srcpixBuf.ScaleSimple(width,width,gdkpixbuf.INTERP_TILES)
	image := gtk.NewImageFromPixbuf(pixBuf)
	return image
}