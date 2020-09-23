package main

import (
	"fmt"
	// "gostart/gohangman/hangman"
	"os"

	"gostart.org/gohangman/hangman"
	// "gostart.go/hangman/hangman"
)

func main() {
	hangman.DrawWelcome()
	// hangman.Draw
	g := hangman.New(8, "golangenforce")
	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Lecture depuis le terminal impossible: %v\n", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
		// fmt.Printf("Lettre proposée: %v\n", strings.ToUpper(l))
	}
	// fmt.Println(g)

}
