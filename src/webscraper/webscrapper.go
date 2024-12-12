package webscraper

import (
  "net/http"
  "fmt"
  "image"
  _"image/jpeg"
  _"image/png"
  _"image/gif"
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
  Src string
  Img image.Image
  Name string
}

func (di *DiscoverItem) FetchImg(src string) {
  di.Img = getImageFromUrl(src)
}

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
  println("DEFAULT SOURCES LEN", len(DEFAULT_SOURCES))
  for _, s := range DEFAULT_SOURCES {
    println("add fetcher")
    fetchers = append(fetchers, SOURCES[s]) 
  }
  w := &Webscraper{
    sources: fetchers,
  }
  return w
} 

func (w* Webscraper) FetchDiscoverItems()  []DiscoverItem {
  println("FetchDiscoverItems")
  var items []DiscoverItem
  for _, f := range w.sources {
    println("fetching")
    fitems := f.FetchDiscoverItems() 
    println("fetched")
    items = append(items, fitems...)
  }
  return items
}


func getImageFromUrl(url string) image.Image {

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
