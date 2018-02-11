package main

import (
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", display)

	http.ListenAndServe(":1337", nil)
}

func display(w http.ResponseWriter, r *http.Request) {
	fp:= gofeed.NewParser()
	
	var url string = "https://pub.scotch.io/feed"
	
	feed, _ := fp.ParseURL(url)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, feed.Title + " - " + feed.Description)
	
	io.WriteString(w, "<!DOCTYPE html><head><link rel='stylesheet' href='//fonts.googleapis.com/css?family=Lato' /><style>body { width: 960px; max-width: 100%; margin: auto; font-family: 'Lato', sans-serif;}.content { display: inline-block; margin-top: 10px; margin-bottom:10px; border: 1px solid #ccc; border-radius: 5px; width: 90%; padding: 2rem; transition: 0.3s all;} .content:hover{ box-shadow: 0px 5px 0px #ccc} a { text-decoration: none; border: 1px solid #ccc; color: #333; padding: 7px; display: inline-block; transition: 0.5s all; } a:hover { background: #ccc; color: #000}</style><title>" + feed.Title + " - " + feed.Description + "</title><body>")
	io.WriteString(w, "<h2>" + feed.Title + "</h2>" + "<p><em>" + feed.Description + "</em></p>")
		
	for i := 0; i<= len(feed.Items)-1; i++ {
		io.WriteString(w, "<div class='content'>")
		io.WriteString(w, "<h2>" + feed.Items[i].Title + "</h2>")

		io.WriteString(w, "<p> Published by " + feed.Items[i].Author.Name + " at " + feed.Items[i].Published + "</p>")

		io.WriteString(w, "<a target='_blank' href=//" + feed.Items[i].Link +">Read More</a>")
    io.WriteString(w, "</div>")

	}
	
	io.WriteString(w, "<h2><a href='"+ feed.Link + "'>Content taken from here</a></h2>")
	io.WriteString(w, "</body></html>")
}