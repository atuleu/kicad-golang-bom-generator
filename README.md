# kicad-golang-bom-generator

A [KiCad](http://www.kicad-pcb.org/) BOM tool generator. It simply
generates a CSV qith grouped entries. An entry is grouped if its
Value, Footprint and any Custom Field are the Same.

## Installation

1. Executable installation:
```bash
go get -u github.com/atuleu/kicad-golang-bon-generator
```

2. Within KiCad
  1. Go to BOM Generator
  2. Click Add Plugin
  3. Select @$GOPATH/bin/kicad-golang-bom-generator@
  4. Edit the command line to replace @<"%I">@ by @"%I"@



## LICENSE

This project is Licensed under GPL version 3


