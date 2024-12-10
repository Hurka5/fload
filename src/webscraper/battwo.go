
package webscraper

import (
  "net/http"
  "fmt"
  "os"
  "log"
  "io/ioutil"
)

type BattwoFetcher struct {
  source string
}

func (f *BattwoFetcher) FetchDiscoverItems() []DiscoverItem {
  println("battwo FetchDiscoverItems")
  var items []DiscoverItem

  resp, err := http.Get(f.source)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", resp.StatusCode)
	}

  // Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

  println(string(body))

  return items
}

func (f *BattwoFetcher) FetchChapters(DiscoverItem) []ChapterItem {
  var items []ChapterItem
  return items
}
func (f *BattwoFetcher) FetchChapter(ChapterItem) {}
