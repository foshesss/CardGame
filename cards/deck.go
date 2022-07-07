package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create and hold all methods of type 'deck', a slice of strings.

type deck []string

// -- Constants
var SUITS = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var VALUES = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

func applyFilePath(fileName string) string {
	return fmt.Sprintf("dataStore/%s", fileName)
}

func newDeck() deck {
	// deck 'constructor', default/standard
	cards := deck{}

	for _, suit := range SUITS {
		for _, value := range VALUES {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	// deck 'constructor', using file
	fileName = applyFilePath(fileName)
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Printf("[newDeckFromFile] Error reading file '%s'\n", fileName)

		if bs == nil {
			return newDeck()
		}

		os.Exit(1)
	}

	str := string(bs)
	return deck(strings.Split(str, ","))
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) deal(numCards int) (deck, deck) {
	return d[numCards:], d[:numCards] // deck new value, hand
}

func (d deck) saveToFile(fileName string) error {
	fileName = applyFilePath(fileName)
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
