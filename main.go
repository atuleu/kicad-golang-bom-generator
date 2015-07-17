package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
)

func Execute() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Syntax: kicad-go-bom-generator <input-xml> <output>")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	if len(path.Ext(outputFile)) == 0 {
		outputFile = outputFile + ".csv"
	}

	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	dec := xml.NewDecoder(input)
	var e KicadExport
	err = dec.Decode(&e)
	if err != nil {
		return err
	}

	BOM, err := NewBOM(e)
	if err != nil {
		return err
	}

	BOM.WriteCSV(output)
	log.Printf("Created BOM in '%s'", outputFile)
	return nil
}

func main() {
	log.SetFlags(0)
	if err := Execute(); err != nil {
		log.Fatalf("Exited after error: %s", err)
	}
}
