package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for i := 0; i <= 100; i += 1 {
		increment(i)
	}
	time.Sleep(time.Millisecond * 350) // Take a beat
	fmt.Print("\033[H\033[2J")         // Clear console
	fmt.Println("Download Finished!")
}

func increment(iterator int) {
	fmt.Print("\033[H\033[2J")                              // Clear console
	fmt.Printf("Downloading (%d/100) bytes...\n", iterator) // Print progress bar. Iterator is current percent
	fmt.Println("[", strings.Repeat("█", iterator), "]")
	time.Sleep(time.Millisecond * 200)
}
