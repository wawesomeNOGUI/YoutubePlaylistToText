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
var doneTraversing = false

func traverseHTML(n *html.Node) {
	if doneTraversing {
		return
	}

	if n.Type == html.TextNode {
		// remove whitespace
		tmp := strings.ReplaceAll(n.Data, " ", "")
		tmp = strings.ReplaceAll(tmp, "\n", "")
		tmp = strings.ReplaceAll(tmp, "\t", "")

		if tmp != "" {
			if tmp == "Suggestions" {	// ignore youtube playlist suggestions at bottom
				doneTraversing = true
				return
			}
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

	var tmpSlice []string
	var totalPlaylistLengthMin int
	var totalPlaylistLengthSec int

	// get song title | artist name
	for i := 0; i < len(songSlice); i++ {
		tmpSlice = append(tmpSlice, songSlice[i] + " | " + songSlice[i + 1])

		// find where next song starts using time as delimiter
		for j := i; j < len(songSlice); j++ {
			if !strings.Contains(songSlice[j], ":") {
				continue
			}

			// try to get min:sec
			time := strings.Split(songSlice[j], ":")

			m, err1 := strconv.Atoi(time[0])
			s, err2 := strconv.Atoi(time[1])

			if err1 == nil && err2 == nil {
				totalPlaylistLengthMin += m
				totalPlaylistLengthSec += s
				
				i = j
				break
			}		
		}
	}

	// convert times
	totalPlaylistLengthMin += totalPlaylistLengthSec / 60
	totalPlaylistLengthSec = totalPlaylistLengthSec % 60

	totalPlaylistLengthHour := totalPlaylistLengthMin / 60
	totalPlaylistLengthMin = totalPlaylistLengthMin % 60

	// print out playlist length and then all the songs
	fmt.Printf("Playlist Length: %d songs!\n\n", len(tmpSlice))

	fmt.Printf("Playlist Length: %d Hours : %d Minutes : %d Seconds!\n\n", totalPlaylistLengthHour, totalPlaylistLengthMin, totalPlaylistLengthSec)
	
	for _, v := range tmpSlice {
		fmt.Println(v)
	}

}
