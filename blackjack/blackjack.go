package blackjack

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

// ErrTableFull the table is full
var ErrTableFull = errors.New("Table is full")

// ErrNameExists name exists already
var ErrNameExists = errors.New("Name already exists")

// ErrNoPlayer no player matches
var ErrNoPlayer = errors.New("Unable to find player")

// Game interface to Blackjack
type Game struct {
	dealer  *dealer
	players map[string]*player
	decks   []*deck
}

type dealer struct {
	hand hand
}

type player struct {
	uuid uuid.UUID
	name string
	hand hand
}

type hand struct {
	cards []card
}

type card struct {
	suit suit
	rank rank
}

type deck struct {
	cards []card
}

type suit = uint8
type rank = uint8

const (
	club suit = iota + 1
	diamond
	heart
	spade
)

const (
	deuce rank = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten   = 10
	jack  = 10
	queen = 10
	king  = 10
	ace   = 11
)

var suits = [4]suit{club, diamond, heart, spade}
var ranks = [13]rank{
	deuce,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
	ten,
	jack,
	queen,
	king,
	ace,
}

// NewGame returns new game instance
func NewGame() *Game {
	return &Game{
		dealer:  &dealer{},
		players: map[string]*player{},
		decks:   []*deck{newDeck()},
	}
}

// FirstDeal give 2 cards to each player
func (g *Game) FirstDeal() error {
	return nil
}

// AddPlayer add a new player to the game
func (g *Game) AddPlayer(name string) error {
	if len(g.players) >= 7 {
		return ErrTableFull
	}

	if _, ok := g.players[name]; ok {
		return ErrNameExists
	}

	newPlayer := newPlayer(name)
	g.players[name] = newPlayer
	return nil
}

// RemovePlayer remove an existing player from the game
func (g *Game) RemovePlayer(name string) error {
	if _, ok := g.players[name]; !ok {
		return ErrNoPlayer
	}

	delete(g.players, name)

	return nil
}

// ResetHands resets dealer and player hands to be empty
func (g *Game) ResetHands() {
	g.dealer.hand = emptyHand()
	for _, player := range g.players {
		player.hand = emptyHand()
	}
}

// Shuffle re-orders the deck(s)
func (g *Game) Shuffle() {
	for _, d := range g.decks {
		d.shuffle()
	}
}

/********
* PRIVATE
********/

func newPlayer(name string) *player {
	return &player{
		name: name,
		uuid: uuid.New(),
		hand: emptyHand(),
	}
}

func newDeck() *deck {
	return &deck{
		cards: newCards(),
	}
}

func newCards() []card {
	cards := []card{}

	for _, s := range suits {
		for _, r := range ranks {
			cards = append(cards, card{s, r})
		}
	}

	return cards
}

func (d *deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func emptyHand() hand {
	return hand{cards: []card{}}
}
