package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/dalloriam/facebook-extractor/facebook"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter partial name: ")
	fname, _ := reader.ReadString('\n')
	_, err := facebook.NewArchive(strings.TrimSpace(fname))

	if err != nil {
		fmt.Println(err)
	}
}
