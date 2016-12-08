package facebook

import (
	"os"
	"strings"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Archive represents an archive of facebook conversations
type Archive struct {
	Owner   string    `json:"owner"`
	Threads []*Thread `json:"threads"`
}

// NewArchive creates a facebook archive by parsing the htm file passed as parameter
func NewArchive(archivePath string, search string) (*Archive, error) {
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

	threads := []*Thread{}

	for arc := range threadChan {
		for _, part := range arc.Participants {
			if search != "" {
				if strings.Contains(part, search) {
					threads = append(threads, arc)
					break
				}
			} else {
				threads = append(threads, arc)
				break
			}
		}
	}

	return &Archive{Owner: owner, Threads: threads}, nil
}
