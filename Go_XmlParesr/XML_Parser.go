package main

import (
    "io/ioutil"
    "encoding/xml"
    "fmt"
)

type opml struct {
    XMLName     xml.Name    `xml:"opml"`
    Version     string      `xml:"version,attr"`
    Body        body
}

type body struct {
    XMLName     xml.Name    `xml:"resources"`
    Outlines    []outline   `xml:"string"`
}

type outline struct {
    Name        string      `xml:"name,attr"`
    Value       string      `xml:",chardata"`
}

func main() {
    o := opml{}
/*
    xmlContent := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<opml version="1.0">
<resources>
  <string name="Hacker News" > aaaa aa </string>
</resources>
</opml>`)
*/
    xmlContent, _ := ioutil.ReadFile("strings_language.xml")
    headXml := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<opml version="1.0">`)
    tailXml := []byte(`</opml>`)
    newXml := append(headXml, xmlContent...)
    newXml = append( newXml,tailXml...)

    err := xml.Unmarshal(newXml, &o)
    if err != nil { panic(err) }
    for _, outline := range o.Body.Outlines {
        fmt.Println(outline.Name+":",outline.Value)

    }
}
