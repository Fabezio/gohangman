package main

import (
	"fmt"
	"os"
	"strings"

	"gostart.go/hangman/hangman"
)

func main() {
	g := hangman.New(8, "golang")
	fmt.Println(g)
	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Lecture depuis le terminal impossible: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Lettre proposée: %v\n", strings.ToUpper(l))
}
