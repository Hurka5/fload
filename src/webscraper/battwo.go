
package webscraper

import (
  "net/http"
  "fmt"
  "os"
  "log"
  _"io/ioutil"
  "golang.org/x/net/html"
)

type BattwoFetcher struct {
  source string
}

func findElementByID(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return n
			}
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if result := findElementByID(child, id); result != nil {
			return result
		}
	}
	return nil
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

  doc, err := html.Parse(resp.Body)
  if err != nil {
		log.Fatalf("Error: %d", err)
    return nil
  }

  println("find element by id")
  slist := findElementByID(doc, "series-list")

  println(slist.Data)
  for child := slist.FirstChild; child != nil; child = child.NextSibling {
    println(child.Data)
    for _,v := range child.Attr {
      println(v.Key," = ",v.Val)
    }
    println(child.Attr)
	}

  return items
}

func (f *BattwoFetcher) FetchChapters(DiscoverItem) []ChapterItem {
  var items []ChapterItem
  return items
}
func (f *BattwoFetcher) FetchChapter(ChapterItem) {}
