package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := flag.String("url", "", "url of youtube playlist")
	file := flag.String("file", "", "path to file from youtube view-source (or just file name if the file is in this directory)")
	flag.Parse()

	if *url == "" && *file == "" {
		fmt.Println("Usage example: main.exe -url https://www.youtube.com/playlist?list=PL96C35uN7xGLLeET0dOWaKHkAlPsrkcha")
		fmt.Println("Usage example: main.exe -file myMusic.html")
		return
	}

	var b []byte //Stores byte text data
	var err error

	if *file != "" {
		b, err = ioutil.ReadFile(*file)
		if err != nil {
			panic(err)
		}
	} else {
		resp, err := http.Get(*url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	}

	//fmt.Println(body)

	var vidTitles []string
	match := "title=\""
	//var match string = "\"title\":{\"runs\":[{\"text\":\""
	//var match2 string = "\"accessibility\":{\"accessibilityData\":{\"label\":\""
	//var match2 string = "aria-label=\""

	body := string(b) //copy of body to work on

	// update april 2024
	// looks like the format is title="song-title" .... .... title="artist-name/channel name"
	// so I'll try to format like that when printing out vidTitles
	for {
		i := strings.Index(body, match)
		if i != -1 {
			//delete everything before that point in the string body
			body = body[i+len(match):]

			titleEnd := strings.Index(body, "\"")
			if titleEnd == -1 {
				break
			}

			//else add title to vidTitles
			vidTitles = append(vidTitles, body[:titleEnd])

			//finally delete that title from main string
			body = body[titleEnd:]

		} else {
			break
		}
	}

	//First Print Playlist Length
	fmt.Println(len(vidTitles) / 2, " Items In Playlist")

	// body = string(b) //fresh copy of body

	// match2 := "aria-label=\""

	// for {
	// 	i := strings.Index(body, match2)
	// 	if i != -1 {
	// 		//delete everything before that point in the string body
	// 		body = body[i+len(match2):]

	// 		titleEnd := strings.Index(body, "\"")
	// 		if titleEnd == -1 {
	// 			break
	// 		}

	// 		//else add title to vidTitles
	// 		vidTitles = append(vidTitles, body[:titleEnd])

	// 		//finally delete that title from main string
	// 		body = body[titleEnd:]
	// 	} else {
	// 		break
	// 	}
	// }

	//Print all the titles
	for i := 0; i < len(vidTitles); i += 2 {
		// list song name : artist name
		fmt.Println(vidTitles[i], ": ", vidTitles[i+1])
	}

}
