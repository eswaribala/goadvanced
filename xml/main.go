package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	XMLName  xml.Name `xml:"notes"`
	NoteList []Note   `xml:"note"`
}

type Note struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("notes.xml")
	//fmt.Println(len(data))
	var notes Notes

	_ = xml.Unmarshal([]byte(data), &notes)
	fmt.Println(len(notes.NoteList))
	for i := 0; i < len(notes.NoteList); i++ {
		fmt.Println(notes.NoteList[0].To)
		fmt.Println(notes.NoteList[i].From)
		fmt.Println(notes.NoteList[i].Heading)
		fmt.Println(notes.NoteList[i].Body)
	}
}
