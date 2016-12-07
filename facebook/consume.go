package facebook

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// ThreadConsumer reads a threadnode from a selection, parses it, and writes a Thread struct
func ThreadConsumer(owner string, inChan <-chan *goquery.Selection, outChan chan<- *Thread, wg *sync.WaitGroup) {
	for threadNode := range inChan {
		threadNames := strings.Split(threadNode.Children().Remove().End().Text(), ", ")
		threadNamesClean := []string{}

		for _, name := range threadNames {
			if name != owner {
				threadNamesClean = append(threadNamesClean, name)
			}
		}
		outChan <- &Thread{Participants: threadNamesClean}
	}
	wg.Done()
}
