package shulamah

import (
	"log"
	"os"
)

func FileReader() string{
	// Read the original file
	data, err := os.ReadFile("sample.txt") //your actual file path
	if err != nil {
		log.Fatal(err)
	}

	// Convert the content to a string
	content := string(data)

	return content
}
