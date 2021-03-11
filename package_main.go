package main

import (
	blackjack "Grubalub/blackjack_ai"
	"fmt"
)

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           1,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
