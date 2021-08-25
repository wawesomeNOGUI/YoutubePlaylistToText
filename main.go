package main

import (
  "fmt"
  "net/http"
  "io"
  "strings"
  "flag"
)

func main(){
  url := flag.String("url", "", "url of youtube playlist")
	flag.Parse()

  if *url == "" {
		fmt.Println("Usage example: main.exe -url https://www.youtube.com/playlist?list=PL96C35uN7xGLLeET0dOWaKHkAlPsrkcha")
		return
	}


  resp, err := http.Get(*url)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  b, err := io.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  body := string(b)

  //fmt.Println(body)

  var vidTitles []string
  var match string = "\"title\":{\"runs\":[{\"text\":\""


  for{
    i := strings.Index(body, match)
    if i == -1 {break}

    //delete everything before that point in the string body
    body = body[i+len(match):]
/*
    titleStart := strings.Index(body, "title")
    if titleStart == -1 {break}

    titleStart = titleStart + len("title") + 2  //length of word title and ="

    body = body[titleStart:]
*/
    titleEnd := strings.Index(body, "\"")
    if titleEnd == -1 {break}

    //else add title to vidTitles
    vidTitles = append(vidTitles, body[:titleEnd])

    //finally delete that title from main string
    body = body[titleEnd:]


  }

  for i:= 0; i<len(vidTitles); i++ {
    fmt.Println(vidTitles[i])
  }


}
