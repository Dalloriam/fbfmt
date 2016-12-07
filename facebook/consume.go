package facebook

import (
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// ThreadConsumer reads a threadnode from a selection, parses it, and writes a Thread struct to outChan
func ThreadConsumer(owner string, inChan <-chan *goquery.Selection, outChan chan<- *Thread, wg *sync.WaitGroup) {
	for threadNode := range inChan {

		// Get the messages
		messages := []*Message{}

		var sender string
		var meta string

		threadNode.Children().Each(func(i int, msg *goquery.Selection) {
			if msg.Nodes[0].Data == "div" {
				sender = msg.Find(".user").First().Text()
				meta = msg.Find(".meta").First().Text()
			} else if msg.Nodes[0].Data == "p" {
				newMsg := &Message{Sender: sender, Metadata: meta, Body: msg.Text()}
				messages = append(messages, newMsg)
			}
		})

		// Get the list of Participants
		threadNames := strings.Split(threadNode.Children().Remove().End().Text(), ", ")

		outChan <- &Thread{Participants: threadNames, Messages: messages}
	}
	wg.Done()
}
