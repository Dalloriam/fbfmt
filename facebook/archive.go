package facebook

import (
	"os"
	"strings"

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

	doc.Find(".thread").Each(func(i int, s *goquery.Selection) {

	})

	return &Archive{}, nil
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
