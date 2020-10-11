package main

import (
	"fmt"
	"strings"
	"time"
)

const bar = "█" // Unicode character, acts as bar element

func main() {
	barIncrement()
}

func barIncrement() {
	var timeStart = time.Now() // Get the current time

	for i := 0; i <= 100; i += 1 {
		timeElapsed := time.Since(timeStart)             // Store the elapsed time
		fmt.Print("\033[H\033[2J")                       // Clear console
		fmt.Printf("Downloading (%d/100) bytes...\n", i) // Print progress bar. Iterator is current percent
		fmt.Println("[", strings.Repeat(bar, i), "]")    // Print bars, relational to the percentage i
		fmt.Println(timeElapsed)
		time.Sleep(time.Millisecond * 200)
	}

	time.Sleep(time.Millisecond * 350) // Take a beat
	fmt.Print("\033[H\033[2J")         // Clear console
	fmt.Println("Download Finished!")
}
