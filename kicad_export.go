package main

import "encoding/xml"

type ComponentField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

type ComponentFieldList struct {
	List []ComponentField `xml:"field"`
}

type Component struct {
	Reference string             `xml:"ref,attr"`
	Value     string             `xml:"value"`
	Footprint string             `xml:"footprint"`
	Fields    ComponentFieldList `xml:"fields"`
}

type ComponentList struct {
	List []Component `xml:"comp"`
}

type KicadExport struct {
	Version    string        `xml:"version,attr"`
	XMLName    xml.Name      `xml:"export"`
	Components ComponentList `xml:"components"`
}
