package main

import (
  "math"
  "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
  "fload/internal/webscraper"
  "fload/internal/gui"
  "fload/pkg/gridwraplist"
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
  // Get items
  items := wb.FetchDiscoverItems()
  
  // Put them in a grids
	/*grid := container.New(layout.NewGridLayout(2))
  for _, i := range items {
    item := gui.NewDiscoverItem(i.Name, i.ImgSrc)
    grid.Add(item)
  }*/

  var pairs []fyne.CanvasObject
  for i := 0; i < len(items); i +=2 {
    println(i)
    item := gui.NewDiscoverItem(items[i].Name, items[i].ImgSrc)
    item2 := gui.NewDiscoverItem(items[i+1].Name, items[i+1].ImgSrc)
    pair := container.New(layout.NewHBoxLayout(), item, item2)
    pairs = append(pairs, pair)
  }

  // Put it in a list
  list := widget.NewGridWrapList(
		func() int {
			return int(math.Floor( float64(len(items)/2) ))
		},
		func() fyne.CanvasObject {    
      item := gui.NewDiscoverItem("Loading...", "")
      item2 := gui.NewDiscoverItem("Loading...", "")
      pair := container.New(layout.NewHBoxLayout(), item, item2)
			return pair
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
      // Set titles and load images
      o.(*fyne.Container).Objects[0].(*gui.DiscoverItem).Title.SetText(items[(i*2)].Name)
      if(items[(i*2)].Img == nil) {
        items[(i*2)].Img = webscraper.GetImageFromUrl(items[(i*2)].ImgSrc)
      }
      o.(*fyne.Container).Objects[0].(*gui.DiscoverItem).Img.Image = items[(i*2)].Img
      
      // if odd
      if(len(items) > (i*2)+1){
        o.(*fyne.Container).Objects[1].(*gui.DiscoverItem).Title.SetText(items[(i*2)+1].Name)
        
        if(items[(i*2)+1].Img == nil) {
          items[(i*2)+1].Img = webscraper.GetImageFromUrl(items[(i*2)+1].ImgSrc)
        }
        o.(*fyne.Container).Objects[1].(*gui.DiscoverItem).Img.Image = items[(i*2)+1].Img
      }
		})

  //return container.NewVScroll(grid)
  return list
}
