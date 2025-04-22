package shulamah

import (
	"log"
	"os"
)

func Test() {
	test_1 := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	err := os.WriteFile("sample.txt", []byte(test_1), 0644)
	if err != nil {
		log.Fatal(err)
	}
}