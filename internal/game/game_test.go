package game_test

import (
	pocgame "github.com/gkalogeitonas/poker-odds-calculator/internal/game"
	"github.com/mattlangl/gophe"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	Describe("NewGameForState", func() {
		var state pocgame.GameKnownState

		BeforeEach(func() {
			state = pocgame.GameKnownState{
				OwnCards: []gophe.Card{
					gophe.NewCard("As"),
					gophe.NewCard("Ac"),
				},
				TableCards: []gophe.Card{
					gophe.NewCard("5s"),
					gophe.NewCard("4h"),
					gophe.NewCard("2s"),
				},
				NumberOfPlayers: 3,
			}
		})

		It("initializes a game respecting the initial state", func() {
			game := pocgame.NewGameForState(state)
			Expect(len(game.Players)).To(Equal(3))
			ourHand := game.Players[0].Hand

			Expect(ourHand).To(ContainElements(
				gophe.NewCard("As"),
				gophe.NewCard("Ac"),
			))

			Expect(game.TableCards).To(ContainElements(
				gophe.NewCard("5s"),
				gophe.NewCard("4h"),
				gophe.NewCard("2s"),
			))

			Expect(len(game.Deck)).To(Equal(47))
			Expect(game.Deck).ToNot(ContainElements(
				gophe.NewCard("As"),
				gophe.NewCard("Ac"),
				gophe.NewCard("5s"),
				gophe.NewCard("4h"),
				gophe.NewCard("2s"),
			))
		})
	})

	Describe("Result", func() {
		When("we have the strongest hand", func() {
			It("returns WIN", func() {

			})
		})

		When("we don't have the strongest hand", func() {
			It("returns LOSS", func() {

			})
		})
		When("we share the strongest hand", func() {
			It("returns DRAW", func() {

			})
		})
	})
})
