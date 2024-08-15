// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
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

func main() {
	urls := [8]string{"https://www.agi.it/cronaca/rss", "https://www.agi.it/economia/rss", "https://www.agi.it/politica/rss", "https://www.agi.it/estero/rss", "https://www.agi.it/cultura/rss", "https://www.agi.it/sport/rss", "https://www.agi.it/innovazione/rss", "https://www.agi.it/lifestyle/rss"}

	var i int

	fmt.Println("Agi rss number")
	fmt.Println("0 to cronaca")
	fmt.Println("1 to economia")
	fmt.Println("2 to politica")
	fmt.Println("3 to estero")
	fmt.Println("4 to cultura")
	fmt.Println("5 to sport")
	fmt.Println("6 to innovazione")
	fmt.Println("7 to lifestyle")
	fmt.Print("Select number rss: ")
	fmt.Scan(&i)

	url := urls[i]

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("Error: request timed out after 3 seconds")
		} else {
			fmt.Println("Error getting response:", err)
		}
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
		fmt.Println("Description:", item.Desc)
		fmt.Println("PubDate:", item.PubDate)
		fmt.Println()
	}
}
