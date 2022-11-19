package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mattlangl/gophe"
)

const (
	WIN = iota
	LOSS
	DRAW
)

var DeckCards []gophe.Card

func init() {
	rand.Seed(time.Now().UnixNano())
	initializeDeck()
}

func initializeDeck() {
	suites := [4]byte{'c', 'd', 'h', 's'}
	ranks := [13]byte{
		'2', '3', '4', '5', '6', '7', '8',
		'9', 'T', 'J', 'Q', 'K', 'A',
	}

	for _, suite := range suites {
		for _, rank := range ranks {
			card := gophe.NewCard(fmt.Sprintf("%c%c", rank, suite))
			DeckCards = append(DeckCards, card)
		}
	}
}

type GameKnownState struct {
	OwnCards        []gophe.Card
	TableCards      []gophe.Card
	NumberOfPlayers int
}

type Hand []gophe.Card

type Player struct {
	Hand Hand
}

type Game struct {
	Deck       []gophe.Card
	Players    []Player
	TableCards []gophe.Card
}

func NewGameForState(knownData GameKnownState) Game {
	deck := deckMinusCards(append(knownData.OwnCards, knownData.TableCards...))
	me := Player{}
	for _, c := range knownData.OwnCards {
		me.Hand = append(me.Hand, c)
	}

	players := []Player{me}
	for i := 1; i < knownData.NumberOfPlayers; i++ {
		players = append(players, Player{})
	}

	return Game{
		Deck:       deck,
		TableCards: knownData.TableCards,
		Players:    players,
	}
}

// Finish deals cards until every table has 5 cards and each player
// has 2 cards.
func (g Game) Finish() {
}

func (g Game) Result() int {
	return WIN
}

func deckMinusCards(excludedCards []gophe.Card) []gophe.Card {
	result := []gophe.Card{}

	lookupHash := map[gophe.Card]interface{}{}
	for _, c := range excludedCards {
		lookupHash[c] = nil
	}

	for _, c := range DeckCards {
		if _, ok := lookupHash[c]; !ok {
			result = append(result, c)
		}
	}

	// https://yourbasic.org/golang/shuffle-slice-array/
	rand.Shuffle(
		len(result),
		func(i, j int) { result[i], result[j] = result[j], result[i] },
	)

	return result
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
