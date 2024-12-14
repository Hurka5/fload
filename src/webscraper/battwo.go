
package webscraper

import (
  "net/http"
  "fmt"
  "log"
  _"io/ioutil"
  "github.com/PuerkitoBio/goquery"
)

type BattwoFetcher struct {
  source string
}


func (f *BattwoFetcher) FetchDiscoverItems() []DiscoverItem {
  var items []DiscoverItem

  // Send GET request to website
  resp, err := http.Get(f.source)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}
  defer resp.Body.Close()

  // Check if succeseded
  if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}

  // PArse it 
  doc, err := goquery.NewDocumentFromReader(resp.Body)
  if err != nil {
		log.Fatalf("Error: %d", err)
  }

  doc.Find("#series-list .item").Each(func(i int, s *goquery.Selection) {
		// Get img source
		img_src, _ := s.Find("a img").Attr("src")
		src, _ := s.Find("a").Attr("href")
		title := s.Find(".item-text .item-title").Text()
    item := DiscoverItem{Src: src, Name: title, ImgSrc: img_src, Fetcher: f}
    println("Fetching: ", title)
    items = append(items, item)
	})

  return items
}

func (f *BattwoFetcher) FetchChapters(DiscoverItem) []ChapterItem {
  var items []ChapterItem
  return items
}
func (f *BattwoFetcher) FetchChapter(ChapterItem) {}
