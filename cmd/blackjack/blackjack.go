package main

import (
	"fmt"

	bj "github.com/jonwho/casinobot/blackjack"
)

func main() {
	fmt.Println("Start test here")

	game := bj.NewGame()
	if err := game.AddPlayer("jon"); err != nil {
		fmt.Println(err)
	}
	if err := game.RemovePlayer("bob"); err != nil {
		fmt.Println(err)
	}
}
