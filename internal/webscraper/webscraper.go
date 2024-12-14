package webscraper

import (
  "net/http"
  "fmt"
  "image"
  _"image/jpeg"
  _"image/png"
  _"image/gif"
  "fyne.io/fyne/v2/canvas"
  _"fyne.io/fyne/v2/container"
)

var SOURCES = map[string]ItemFetcher{
  "https://battwo.com": &BattwoFetcher{source: "https://battwo.com"},
}


//TODO: Source management
var DEFAULT_SOURCES = [...]string {
  "https://battwo.com",
}

type DiscoverItem struct {
  Fetcher ItemFetcher
  Object *canvas.Image
  Src string
  ImgSrc string 
  Img image.Image // cache
  Name string
}
/*
func (di *DiscoverItem) FetchImg(obj *canvas.Image) {
  if(di.Img == nil) {
    di.Img = GetImageFromUrl(di.ImgSrc)
  }
 // obj.image = ob.NewMax(di.Img)
}
*/
type ChapterItem struct {
  Src string
  Name string
}

type ItemFetcher interface {
  FetchDiscoverItems() []DiscoverItem       /* Fetches all the discover page items from all the sources */ 
  FetchChapters(DiscoverItem) []ChapterItem  /* Fetches all the chapters and their sources */ 
  FetchChapter(ChapterItem)   /* Fetches all images of the chapter */ 
}

type Webscraper struct {
  sources []ItemFetcher
}

func New() *Webscraper {
  var fetchers []ItemFetcher
  for _, s := range DEFAULT_SOURCES {
    fetchers = append(fetchers, SOURCES[s]) 
  }
  w := &Webscraper{
    sources: fetchers,
  }
  return w
} 

func (w* Webscraper) FetchDiscoverItems()  []DiscoverItem {
  var items []DiscoverItem
  for _, f := range w.sources {
    fitems := f.FetchDiscoverItems() 
    items = append(items, fitems...)
  }
  return items
}


func GetImageFromUrl(url string) image.Image {

  resp, err := http.Get(url)
  if err != nil {
    fmt.Errorf("error making http request: %w", err)
  }
  defer resp.Body.Close()

  img, format, err := image.Decode(resp.Body)
  if err != nil {
    fmt.Errorf("cannot decode image: %w", err)
  }
  fmt.Println("Image format:", format)

  return img
}
