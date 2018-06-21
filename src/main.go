package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("/workspaces/AndriiAndrus/src/mobiSec/text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("text to append\n"); err != nil {
		log.Println(err)
	}
}
