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

  //fmt.Println(body)

  var vidTitles []string
  var match string = "\"title\":{\"runs\":[{\"text\":\""
  var match2 string = "\"accessibility\":{\"accessibilityData\":{\"label\":\""
  //var match string = "aria-label"

  body := string(b)  //copy of body to work on

  for{
    i := strings.Index(body, match)
    if i != -1 {
      //delete everything before that point in the string body
      body = body[i+len(match):]

      titleEnd := strings.Index(body, "\"")
      if titleEnd == -1 {break}

      //else add title to vidTitles
      vidTitles = append(vidTitles, body[:titleEnd])

      //finally delete that title from main string
      body = body[titleEnd:]

    }else{
      break
    }



/*
    titleStart := strings.Index(body, "title")
    if titleStart == -1 {break}

    titleStart = titleStart + len("title") + 2  //length of word title and ="

    body = body[titleStart:]
*/

  //  fmt.Println(body, "\n")

  }

  body = string(b) //fresh copy of body

  for {
    i := strings.Index(body, match2)
    if i != -1 {
      //delete everything before that point in the string body
      body = body[i+len(match2):]

      titleEnd := strings.Index(body, "\"")
      if titleEnd == -1 {break}

      //else add title to vidTitles
      vidTitles = append(vidTitles, body[:titleEnd])

      //finally delete that title from main string
      body = body[titleEnd:]
    }else{
      break
    }
  }



  for i:= 0; i<len(vidTitles); i++ {
    fmt.Println(vidTitles[i])
  }


}
