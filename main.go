// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Desc    string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func without_tags(text string) string {
	newText := ""
	newText = strings.ReplaceAll(text, "<p>", "")
	newText = strings.ReplaceAll(newText, "</p>", "")
	newText = strings.ReplaceAll(newText, "<div>", "")
	newText = strings.ReplaceAll(newText, "</div>", "")
	newText = strings.ReplaceAll(newText, "<strong>", "")
	newText = strings.ReplaceAll(newText, "</strong>", "")
	newText = strings.ReplaceAll(newText, "<h2>", "")
	newText = strings.ReplaceAll(newText, "</h2>", "")
	newText = strings.ReplaceAll(newText, "<h3>", "")
	newText = strings.ReplaceAll(newText, "</h3>", "")
	newText = strings.ReplaceAll(newText, "<br>", "")
	newText = strings.ReplaceAll(newText, "<;>", "")
	newText = strings.ReplaceAll(newText, "<.;>", ".")
	newText = strings.ReplaceAll(newText, "&nbsp;", "")
	return newText
}

func main() {
	urls := [8]string{
		"https://www.agi.it/cronaca/rss",
		"https://www.agi.it/economia/rss",
		"https://www.agi.it/politica/rss",
		"https://www.agi.it/estero/rss",
		"https://www.agi.it/cultura/rss",
		"https://www.agi.it/sport/rss",
		"https://www.agi.it/innovazione/rss",
		"https://www.agi.it/lifestyle/rss",
	}

	i := -1

	fmt.Println("Agi rss number (0 for exit)")
	fmt.Println("1 to cronaca")
	fmt.Println("2 to economia")
	fmt.Println("3 to politica")
	fmt.Println("4 to estero")
	fmt.Println("5 to cultura")
	fmt.Println("6 to sport")
	fmt.Println("7 to innovazione")
	fmt.Println("8 to lifestyle")
	fmt.Print("Select number rss: ")
	fmt.Scan(&i)

	if i == 0 {
		os.Exit(0)
	}

	url := urls[i-1]

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	defer resp.Body.Close()

	var rss Rss
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	fmt.Println("Title:", rss.Channel.Title)
	fmt.Println("Description:", rss.Channel.Description)

	for _, item := range rss.Channel.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Link:", item.Link)
		fmt.Print("Description: ")
		fmt.Println(without_tags(item.Desc))
		fmt.Println("PubDate:", item.PubDate)
		fmt.Println()
	}
}
