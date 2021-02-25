package main

import blackjack "Grubalub/blackjack_ai"

func main() {

	game := blackjack.New()
	game.Play(blackjack.HumanAI())
}
