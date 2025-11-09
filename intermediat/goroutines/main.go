package main

import(
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println("Letter:", string(ch))
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Finished executing goroutines")
}