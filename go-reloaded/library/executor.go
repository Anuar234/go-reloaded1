// package shulamah

// func Executor() string{
// 	text := Bin()
// 	text2 := Hex()
// 	return text2
// }

package shulamah

import (
	"fmt"
	"log"
	"os"
)

func Executor() {
	text1 := Hex(FileReader())
	text2 := Bin(text1)
	text3 := Cap(text2)
	text4 := Low(text3)
	text5 := Up(text4)

	err := os.WriteFile("result.txt", []byte(text5), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully changed sample.txt into result.txt")
}
