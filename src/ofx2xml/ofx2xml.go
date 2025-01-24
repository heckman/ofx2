package main

// remarshals ofx file to at least version 2.0, which is XML

import (
	"fmt"
	"os"

	"github.com/aclindsa/ofxgo"
)

func printOFXFileAsXML(filename string) error {
	// Open the specified OFX file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	// Parse OFX data using ofxgo library
	response, err := ofxgo.ParseResponse(file)
	if err != nil {
		return fmt.Errorf("error parsing OFX file %s: %v", filename, err)
	}

	// Convert to OFX 2.0 for XML if current version only supports SGML
	if response.Version < ofxgo.OfxVersion200 {
		response.Version.FromString("200")
	}

	// Convert parsed response to XML format
	xmlData, err := response.Marshal()
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return fmt.Errorf("error parsing OFX file %s: %v", filename, err)
	}
	// print it
	fmt.Println(xmlData)

	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ofx2xml <ofx_file>")
		os.Exit(1)
	}

	if err := printOFXFileAsXML(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
