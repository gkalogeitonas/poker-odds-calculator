package game

import "github.com/mattlangl/gophe"

const (
	WIN = iota
	LOSS
	DRAW
)

type GameState struct {
	OwnCards   []gophe.Card
	TableCards []gophe.Card
}
type Game struct{}

func NewGameForState(state GameState) Game {
	return Game{}
}

func (g Game) Result() int {
	return WIN
}

// cc := gophe.NewHand(
// 	gophe.NewCard("5c"),
// 	gophe.NewCard("5h"),
// 	gophe.NewCard("2s"),
// 	gophe.NewCard("2h"),
// )
// r := gophe.EvaluateHand(*cc)
// v := r.GetValue()
// fmt.Printf("v = %+v\n", v)
// fmt.Println(r.DescribeCategory())
