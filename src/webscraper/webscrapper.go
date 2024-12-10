package webscraper

var SOURCES = map[string]ItemFetcher{
  "https://battwo.com": &BattwoFetcher{source: "https://battwo.com"},
}


//TODO: Source management
var DEFAULT_SOURCES = [...]string {
  "https://battwo.com",
}

type DiscoverItem struct {
  Src string
  Img string
  Name string
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
