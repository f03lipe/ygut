package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	spew "github.com/davecgh/go-spew/spew"
)

type RssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	XMLName       xml.Name         `xml:"channel"`
	Language      string           `xml:"language"`
	Items         []RssChannelItem `xml:"item"`
	LastBuildDate string           `xml:"lastBuildDate"`
}

type RssChannelItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func parseDescription(body string) {
	r := strings.NewReplacer(
		"&lt;", "<",
		"&gt;", ">",
		"<br>", "<br />")

	fmt.Printf("%v", r.Replace(body))
}

func main() {
	file, err := ioutil.ReadFile("menu.xml")
	if err != nil {
		log.Fatalf("WHAT? %s\n", err)
	}

	v := RssFeed{}
	r := strings.NewReplacer(
		"&", "&amp;",
		"<br>", "<br />")
	finalXml := []byte(r.Replace(string(file)))
	if err := xml.Unmarshal(finalXml, &v); err != nil {
		fmt.Printf("what? %+v\n", err)
	}
	spew.Dump(v)

	parseDescription(v.Channel.Items[0].Description)
}
