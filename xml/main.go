package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	XMLName      xml.Name `xml:"users"`
	noteInstance []Note   `xml:"user"`
}

type Note struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("notes.xml")

	notes := &Notes{}

	_ = xml.Unmarshal([]byte(data), &notes)

	for i := 0; i < len(notes.noteInstance); i++ {
		fmt.Println(notes.noteInstance[i].To)
		fmt.Println(notes.noteInstance[i].From)
		fmt.Println(notes.noteInstance[i].Heading)
		fmt.Println(notes.noteInstance[i].Body)
	}
}
