package main

import (
  "fmt"
  "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
  _"fyne.io/fyne/v2/canvas"
	_"image/color"

  "fload/src/webscraper"
)

var wb *webscraper.Webscraper

func main() {
  
  wb = webscraper.New()

	app := app.New()
	w := app.NewWindow("Fload")
  
  discover := container.NewTabItemWithIcon("Discover", theme.SearchReplaceIcon(), discoverPage())

	tabs := container.NewAppTabs(
	  container.NewTabItemWithIcon("Favorites", theme.ListIcon(), widget.NewLabel("Favorites")),
    discover,
	  container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings")),
	)

  println(discover.Content.Visible())

	tabs.SetTabLocation(container.TabLocationBottom)
  w.Resize(fyne.NewSize(450,800))
	w.SetContent(tabs)
	w.ShowAndRun()
}

func discoverPage() fyne.CanvasObject {
  items := wb.FetchDiscoverItems()
  var ditems []*fyne.Container
  
  for _, i := range items {
    img := canvas.NewImageFromImage(i.Img)
    img.FillMode = canvas.ImageFillStretch
    img.SetMinSize(fyne.Size{Height: 300})


    name := widget.NewLabel(i.Name)
    name.Wrapping = fyne.TextTruncate
    name.Truncation = fyne.TextTruncateEllipsis
    name.Wrapping = 2
    name.Alignment = fyne.TextAlignCenter
 
    di := container.NewVBox(
      img,
      name,
    )

    openButton := widget.NewButton("", func(){
      fmt.Println("Image clicked...")
    })

    // this encapsulate the image and the button
    btn_box := container.NewPadded(openButton, di)
    //btn_box := container.NewStack(openButton, di)
    ditems = append(ditems, btn_box)
  }

	grid := container.New(layout.NewGridLayout(2))
  for _, di := range ditems {
    grid.Add(di)
  }
  return container.NewVScroll(grid)
}
