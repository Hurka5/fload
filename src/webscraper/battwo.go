
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
  println("battwo FetchDiscoverItems")
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
    return nil
  }

  println("find element by id")
  doc.Find("#series-list .item").Each(func(i int, s *goquery.Selection) {
		// Get img source
		img_src, _ := s.Find("a img").Attr("src")
		src, _ := s.Find("a").Attr("href")
		title := s.Find(".item-text .item-title").Text()
    img := getImageFromUrl(img_src)
    item := DiscoverItem{Src: src, Img: img, Name: title}
    println(title)
    println(img_src)
    println(src)
    println("---")
    items = append(items, item)
	})

  return items
}

func (f *BattwoFetcher) FetchChapters(DiscoverItem) []ChapterItem {
  var items []ChapterItem
  return items
}
func (f *BattwoFetcher) FetchChapter(ChapterItem) {}
