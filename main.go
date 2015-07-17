package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func Execute() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("Pass file to parse")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	var e KicadExport
	err = dec.Decode(&e)
	if err != nil {
		return err
	}

	bom, err := NewBom(e)
	if err != nil {
		return err
	}

	log.Printf("%+v", bom)

	return nil
}

func main() {
	log.SetFlags(0)
	if err := Execute(); err != nil {
		log.Fatalf("Exited after error: %s", err)
	}
}
