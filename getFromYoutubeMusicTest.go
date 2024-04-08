package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"

	"golang.org/x/net/html"
)

var songSlice []string

func traverseHTML(n *html.Node) {
	if n.Type == html.TextNode {
		// remove whitespace
		tmp := strings.ReplaceAll(n.Data, " ", "")
		tmp = strings.ReplaceAll(tmp, "\n", "")
		tmp = strings.ReplaceAll(tmp, "\t", "")

		if tmp != "" {
			songSlice = append(songSlice, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTML(c)
	}
}

func main() {
	file := flag.String("file", "", "path to file from youtube view-source (or just file name if the file is in this directory)")
	flag.Parse()

	if *file == "" {
		fmt.Println("Usage example: main.exe -file myMusic.html")
		return
	}

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// pass file to html parser
	doc, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	traverseHTML(doc)

	/*
	format is:
		title
		artist
		subtitles
		...
		vid length (eg. 4:00)
	
	so we'll use video length as delimiter
	*/

	// print song title | artist name
	for i := 0; i < len(songSlice); i++ {
		fmt.Println(songSlice[i], "|", songSlice[i + 1])

		// find where next song starts using time as delimiter
		for j := i; j < len(songSlice); j++ {
			tmp := strings.ReplaceAll(songSlice[j], ":", "")

			_, err := strconv.Atoi(tmp)
			if err == nil {
				i = j
				break
			}
		}
	}

}
