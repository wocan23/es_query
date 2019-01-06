package component

import (
	"github.com/mattn/go-gtk/gtk"
)

const(
	NavItemWidth = 80
	NavItemHeight = 100
)

/**
	创建一个带图标文字带按钮
 */
func GenerateNavItem(width int,height int,imagePath string,title string) *gtk.Button{
	button := gtk.NewButton()
	button.SetUSize(NavItemWidth,NavItemHeight)
	button.SetVisible(true)

	image := gtk.NewImageFromFile(imagePath)
	image.SetUSize(NavItemWidth,NavItemWidth)

	label := gtk.NewLabel(title)
	label.SetUSize(NavItemWidth,NavItemHeight-NavItemWidth)

	button.Add(image)
	button.Add(label)

	return button
}