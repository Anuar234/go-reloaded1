package shulamah

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Hex(text string) string{
	// Use regular expression to find hexadecimal numbers (without "0x")
	re := regexp.MustCompile(`(\b\w+)\s+\(hex\)`)
	// Find all binary numbers in the text
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
	// Convert binary to decimal
	parts := re.FindStringSubmatch(match)

	decimal, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			log.Fatal(err)
			}
		// Return the decimal equivalent
		return fmt.Sprintf("%d", decimal)
		})
	return modifiedText

	// // Write the modified content to a new file
	// err := os.WriteFile("hex.txt", []byte(modifiedText), 0644) // Write to a new file
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Print success message
	// fmt.Println("Hexadecimal numbers have been converted to decimal and written to 'hex.txt'.")
}

// Bin reads text from a file, converts binary numbers to decimal, and writes to bin.txt

// Bin reads text from a file, converts binary numbers to decimal, and writes to bin.txt
func Bin(text string) string{
	// Regex to find binary numbers before "(bin)"
	re := regexp.MustCompile(`\b([01]+)\s+\(bin\)`)

	// Replace binary numbers with their decimal equivalent
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match // Return original match if regex fails
		}

		// Convert binary to decimal
		decimal, err := strconv.ParseInt(parts[1], 2, 64)
		if err != nil {
			return match // Return original text if conversion fails
		}

		// Return decimal equivalent as a string
		return fmt.Sprintf("%d", decimal)
	})
	return modifiedText

	// // Write the modified content to bin.txt
	// err := os.WriteFile("bin.txt", []byte(modifiedText), 0644)
	// if err != nil {
	// 	log.Fatal("Error writing to file:", err)
	// }

	// fmt.Println("Binary numbers have been converted to decimal and written to 'bin.txt'.")
}
