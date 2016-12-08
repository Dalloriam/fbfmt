package main

import (
	"strings"

	"encoding/json"

	"io/ioutil"

	"flag"

	"github.com/dalloriam/facebook-extractor/facebook"
)

var inPath string
var outPath string
var searchTerm string

func init() {
	flag.StringVar(&inPath, "In", "messages.htm", "Path to facebook .HTM file")
	flag.StringVar(&outPath, "Out", "out.json", "Output JSON file to create")
	flag.StringVar(&searchTerm, "Search", "", "Search term")
	flag.Parse()
}

func main() {
	archive, err := facebook.NewArchive(strings.TrimSpace(inPath), searchTerm)

	if err != nil {
		panic(err)
	}

	outStr, err := json.Marshal(archive)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(strings.TrimSpace(outPath), outStr, 0644)

	if err != nil {
		panic(err)
	}
}
