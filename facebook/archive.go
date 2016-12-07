package facebook

import (
	"fmt"
	"os"
	"strings"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Archive represents an archive of facebook conversations
type Archive struct {
	Owner   string
	Threads []Thread
}

// NewArchive creates a facebook archive by parsing the htm file passed as parameter
func NewArchive(archivePath string) (*Archive, error) {
	fileReader, err := os.Open(archivePath)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(fileReader)

	if err != nil {
		return nil, err
	}

	owner := strings.TrimSpace(doc.Find("h1").First().Text())

	// Open the thread input channel & the Thread output channel
	nodeChan := make(chan *goquery.Selection)
	threadChan := make(chan *Thread)
	var wg sync.WaitGroup

	// Start the thread consumers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ThreadConsumer(owner, nodeChan, threadChan, &wg)
	}

	// Start the thread producer
	go func() {
		doc.Find(".thread").Each(func(i int, s *goquery.Selection) {
			nodeChan <- s
		})
		close(nodeChan)
	}()

	go func() {
		wg.Wait()
		close(threadChan)
	}()

	for arc := range threadChan {
		fmt.Println(arc.Participants)
	}

	return &Archive{Owner: owner}, nil
}

// GetThreadByUser find all threads with a specific facebook user
// (can be multiple threads bc of group conversations)
func (a *Archive) GetThreadByUser(user string) []*Thread {
	var threads []*Thread

	for _, th := range a.Threads {
		for _, p := range th.Participants {
			if strings.Contains(p, user) {
				threads = append(threads, &th)
			}
		}
	}

	return threads
}
