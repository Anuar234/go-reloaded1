package shulamah

import (
	"regexp"
	"strings"
)

// Cap reads text from file, modifies words before "(cap)", and writes to capitalized.txt
func Cap(text string) string {

	// Regex to match words before "(cap)"
	re := regexp.MustCompile(`(\b\w+)\s+\(cap\)`)
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match // Return original match if no valid group
		}
		word := parts[1] // Extract word before (cap)

		// Capitalize first letter, lowercase the rest
		return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	})

	return modifiedText

	// Write the modified content to capitalized.txt
	// err := os.WriteFile("capitalized.txt", []byte(modifiedText), 0644)
	// if err != nil {
	// 	log.Fatal("Error writing to file:", err)
	// }

	// fmt.Println("Modified text written to capitalized.txt")
}

func Low(text string) string {

	// Modify text: uppercase words before "(up)"
	re := regexp.MustCompile(`(\b\w+)\s+\(low\)`)
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		return strings.ToLower(parts[1])
	})

	return modifiedText

	// Write the modified content to output.txt
	// err := os.WriteFile("lowed.txt", []byte(modifiedText), 0644)
	// if err != nil {
	// 	log.Fatal("Error writing to file:", err)
	// }

	// fmt.Println("Modified text written to lowed.txt")
}

func Up(text string) string {

	// Modify text: uppercase words before "(up)"
	re := regexp.MustCompile(`(\b\w+)\s+\(up\)`)
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		return strings.ToUpper(parts[1])
	})

	return modifiedText

	// Write the modified content to output.txt
	// err := os.WriteFile("uppped.txt", []byte(modifiedText), 0644)
	// if err != nil {
	// 	log.Fatal("Error writing to file:", err)
	// }

	// fmt.Println("Modified text written to upped.txt")
}
