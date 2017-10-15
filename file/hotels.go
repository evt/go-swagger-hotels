package file

import (
	"fmt"
	"io"
	"os"
)

// TestXMLFile is a test XML file with hotel data
var TestXMLFile = os.Getenv("GOPATH") + "/src/go-swagger-hotels/file/hotels.xml"

// GetXMLFileReader opens XML hotels file and returns reader
func GetXMLFileReader() (io.Reader, error) {
	if !fileExists(TestXMLFile) {
		return nil, fmt.Errorf("File %s doesn't exist", TestXMLFile)
	}
	file, err := os.Open(TestXMLFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// fileExists checks if file exists
func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
