package gui

import (
  "image"
  "image/draw"
  "image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
  "fload/internal/webscraper"
)

type DiscoverItem struct {
  widget.BaseWidget
	Title  *widget.Label
	Img    *canvas.Image
}

func NewDiscoverItem(title, img_src string) *DiscoverItem {


  // Creating placeholder for img
  placeholder := image.NewRGBA(image.Rect(0,0,225,500))
  col := color.RGBA{255, 255, 255, 255}
	draw.Draw(placeholder, placeholder.Bounds(), &image.Uniform{col}, image.Point{}, draw.Src)


  // Creating img
  img := canvas.NewImageFromImage(placeholder)
  img.FillMode = canvas.ImageFillStretch
  img.SetMinSize(fyne.Size{Height: 300, Width: 500})
  
  // Create title label
  t := widget.NewLabel(title)
  t.Wrapping = fyne.TextTruncate
  t.Truncation = fyne.TextTruncateEllipsis
  t.Wrapping = 2
  t.Alignment = fyne.TextAlignCenter

	item := &DiscoverItem{
		Img:    img,
		Title:  t,
	}

  item.ExtendBaseWidget(item)

  //go item.fetchImg(img_src)

	return item
}

//Fetch image and load it
func (item *DiscoverItem) fetchImg(img_src string) {
  img := webscraper.GetImageFromUrl(img_src)
  item.Img.Image = img;
  item.Refresh()
}

func (item *DiscoverItem) CreateRenderer() fyne.WidgetRenderer {
  b := container.NewVBox(item.Img, item.Title)
	return widget.NewSimpleRenderer(b)
}


