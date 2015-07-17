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
	var res KicadExport
	err = dec.Decode(&res)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func main() {
	if err := Execute(); err != nil {
		log.Fatalf("Exited after error: %s", err)
	}
}
