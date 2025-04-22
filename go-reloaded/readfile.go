// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("sample.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan(){
// 		fmt.Println(scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	//read the file

// 	data, err := os.ReadFile("sample.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// converting the content to a string and replace the word

// 	content := string(data)

// 	modifiedContent := strings.Replace(content, "niggers", "gentlemen", -1)

// 	// write the modified content back to the file

// 	err = os.WriteFile("sample.txt", []byte(modifiedContent), 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//printing successfull message
// 	fmt.Println("File content updated successfully")
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	//read the file

// 	data, err := os.ReadFile("sample.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// converting the content to a string and replace the word

// 	content := string(data)

// 	modifiedContent := strings.Replace(content, "niggers", "gentlemen", -1)

// 	// write the modified content back to the file

// 	err = os.WriteFile("result.txt", []byte(modifiedContent), 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//printing successfull message
// 	fmt.Println("File content modified successfully")
// }

package main

import (
	shulamah "shulamah/library"
)

func main() {
	shulamah.Test()
	shulamah.Executor()
}
