package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println(`
	          _______  _        _______  _______  _______  _       
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
																  
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}
func drawState(g *Game, guess string) {
	// g.FoundLetters = strings.ToUpper(g.FoundLetters)
	fmt.Print("Lettres Devinées: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Lettres Utilisées: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Bien joué! \n")
	case "alreadyGuessed":
		fmt.Printf("Lettre '%s' déjà tentée \n", guess)
	case "badGuess":
		fmt.Printf("Non, '%s' n'est pas dans ce mot \n", guess)
	case "lost":
		fmt.Print("Perdu! :( Le mot était: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print(`
		  .-""""""-.
		.'          '.
	   /   O      O   \
	  :                :
	  |                |   
	  : ',          ,' :
	   \  '-......-'  /
		'.          .'
	      '-......-'
		`)
		fmt.Print("BRAVO!  Le mot était: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
