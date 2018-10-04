package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// cards = append(cards, "Five of Diamonds")
	// var card string = "Ace of Spades"
	// cards := deck{"Ace of Spades", newCard()}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	sizeCard := r.Intn(5)

	cards := newDeck()
	//deal card
	hand, reaminingCards := deal(cards, sizeCard)
	hand.print()
	reaminingCards.print()

	//save to file
	sizeStr := strconv.Itoa(sizeCard)
	errSaveToFile := cards.saveToFile("my_cards" + sizeStr)
	if errSaveToFile != nil {
		fmt.Println("error read from file", errSaveToFile)
	}

	//read file
	newCards, errFromFile := newDeckFromFile("my_cards")
	if errFromFile != nil {
		fmt.Println("error read from file", errFromFile)
	}
	newCards.print()
	fmt.Println(newCards.toString())

	//shuffle card
	// cards := newDeck()
	cards.shuffle()
	cards.print()
	fmt.Println("=================")
	print()

}
