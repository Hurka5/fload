package main

import (
  "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
  _"fyne.io/fyne/v2/canvas"
	_"image/color"

  "fload/src/webscraper"
)

var wb *webscraper.Webscraper

func main() {
  
  wb = webscraper.New()

	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")



	tabs := container.NewAppTabs(
	  container.NewTabItemWithIcon("Discover", theme.SearchReplaceIcon(), discoverPage()),
	  container.NewTabItemWithIcon("Favorites", theme.ListIcon(), widget.NewLabel("Favorites")),
	  container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func discoverPage() *fyne.Container {
  items := wb.FetchDiscoverItems()
  var ditems []*fyne.Container
  
  for _, i := range items {
    di := container.NewVBox(
      widget.NewLabel(i.Name),
      widget.NewLabel(i.Src),
      widget.NewLabel(i.Img),
    )
    ditems = append(ditems, di)
  }

	grid := container.New(layout.NewGridLayout(2))
  for _, di := range ditems {
    grid.Add(di)
  }
  return grid
}
