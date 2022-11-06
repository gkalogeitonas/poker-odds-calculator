package main

import (
	"fmt"

	"github.com/gkalogeitonas/poker-odds-calculator/internal/game"
	gamepkg "github.com/gkalogeitonas/poker-odds-calculator/internal/game"
	// "github.com/mattlangl/gophe"
)

const MonteCarloIterations = 1000

func main() {
	currentState := game.GameState{}

	wins, losses, draws := 0, 0, 0
	for i := 0; i < MonteCarloIterations; i++ {
		game := gamepkg.NewGameForState(currentState)

		switch result := game.Result(); result {
		case gamepkg.WIN:
			wins++
		case gamepkg.LOSS:
			losses++
		case gamepkg.DRAW:
			draws++
		}
	}
	fmt.Printf("wins = %+v\n", wins)
	fmt.Printf("losses = %+v\n", losses)
	fmt.Printf("draws = %+v\n", draws)

}
