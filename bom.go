package main

import "fmt"

type Bom struct {
	Headers []string
	Items   [][]string
}

func NewBom(e KicadExport) (*Bom, error) {
	return nil, fmt.Errorf("not implemenetd")
}
