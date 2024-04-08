# YoutubePlaylistToText
Output all the video names in a youtube playlist in plain text.

*Note*

To get all the names in a long playlist you have to:
- Scroll all the way to the end of the playlist to get youtube to send you all the video names
- Next go into the dev console and inspect one of the video titles to get the div it’s in
  (The farthest div you should have to go will be labled something like: <div id="content" class="style-scope ytd-app">)
- Copy everything inside that div to get all the video titles
- Paste that text into a file and use that file as input when running the program e.g:
main.exe -file pathToFile.html > playListAsText.txt
