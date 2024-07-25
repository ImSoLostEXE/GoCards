package main

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Deck []Card

type Card struct {
	rank Rank
	suit string
}

type Rank struct {
	cardname string
	value    int
}

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{" of Spades", " of Clubs", " of Diamonds", " of Hearts"}
	cardRanks := []Rank{
		{cardname: "Ace", value: 1},
		{cardname: "Two", value: 2},
		{cardname: "Three", value: 3},
		{cardname: "Four", value: 4},
		{cardname: "Five", value: 5},
		{cardname: "Six", value: 6},
		{cardname: "Seven", value: 7},
		{cardname: "Eight", value: 8},
		{cardname: "Nine", value: 9},
		{cardname: "Ten", value: 10},
		{cardname: "Jack", value: 11},
		{cardname: "Queen", value: 12},
		{cardname: "King", value: 13},
	}

	for _, suit := range cardSuits {
		for _, rank := range cardRanks {
			cards = append(cards, Card{rank, suit})
		}
	}

	return cards
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card.rank.cardname+card.suit, " ", card.rank.value)
	}
}

func byteToInt(b byte) int {
	v, err := strconv.Atoi(string(b))
	if err != nil {
		fmt.Println("failed to convert byte to int")
	}
	return v
}

func (d Deck) toString() string {
	cardStringSlice := []string{}
	for _, card := range d {
		cardStringSlice = append(cardStringSlice, card.rank.cardname+card.suit+strconv.Itoa(card.rank.value))
	}
	return strings.Join([]string(cardStringSlice), ",")
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(filename string) []byte {
	bsl, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return bsl
}

func bytesliceToDeck(d []byte) Deck {
	deck := bytes.Split(d, []byte(","))
	savedDeck := Deck{}

	v := false
	s := false

	for _, card := range deck {
		tmpcardname := ""
		tmpvalue := 0
		tmpsuit := ""
		s = false
		v = false

		for _, character := range card {
			if character == 32 {
				s = true
			} else if character == 115 {
				tmpsuit += string(character)
				s = false
				v = true
			}
			//comparison to bool for readability
			if s == true {
				tmpsuit += string(character)
			} else if v == true {
				if tmpvalue != 0 {
					tmpvalue += 10
				} else {
					tmpvalue += byteToInt(character)
				}
			} else {
				tmpcardname += string(character)
			}
		}
		savedDeck = append(savedDeck, Card{
			rank: Rank{
				cardname: tmpcardname,
				value:    tmpvalue,
			},
			suit: tmpsuit,
		})
	}
	return savedDeck
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//HTML Handlers

type HtmlCard struct {
	cardname string
	suit     string
	value    int
}

func (c Card) htmlFormatCard() HtmlCard {

	htmlcard := HtmlCard{}

	htmlcard.cardname = c.rank.cardname
	htmlcard.suit = c.suit
	htmlcard.value = c.rank.value

	return htmlcard
}

func (c Card) htmlFormatCname() string {

	htmlCname := c.rank.cardname

	return htmlCname
}

func (c Card) htmlFormatCsuit() string {

	htmlCsuit := c.suit

	return htmlCsuit
}

func (c Card) htmlFormatCval() int {

	htmlCval := c.rank.value

	return htmlCval
}

var fm = template.FuncMap{
	"smplc": Card.htmlFormatCard,
	"cname": Card.htmlFormatCname,
	"csuit": Card.htmlFormatCsuit,
	"cval":  Card.htmlFormatCval,
}
