package component

import (
	"github.com/mattn/go-gtk/gtk"
)


/**
	创建一个带图标文字带按钮
 */
func GenerateNavItem(width int,height int,imagePath string,title string) *gtk.Button{
	button := gtk.NewButton()
	button.SetUSize(width,height)
	button.SetVisible(true)

	image := gtk.NewImageFromFile(imagePath)
	image.SetUSize(width,height)

	label := gtk.NewLabel(title)
	label.SetUSize(width,height)

	button.Add(image)
	button.Add(label)

	return button
}