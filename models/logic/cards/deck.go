package main

import(
	"fmt"
	"strings"
	"io/ioutil"	
	"math/rand"
	"time"
)

// Create a new  type of `deck`
// which is a slice of string
type deck []string 

//create deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}
	
	for _, suit := range cardSuits{
		for _,value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}	

//print card
func (d deck) print() {
	for i, card := range d{
		fmt.Println(i,card)
	}
}

//card deal
func deal(d deck,handSize int) (ori deck, reverse deck) {
	ori = d[:handSize]
	reverse = d[handSize:]
	return 
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")	
}

//save to file
func (d deck) saveToFile(filename string) error {	
	return ioutil.WriteFile(filename, []byte(d.toString()),0666)	
}

//get deck from file
func newDeckFromFile(filename string) (deck, error) {
	bs, err :=ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ",err)	
		return nil, err	
	}
	s := strings.Split(string(bs), ",")
	return deck(s), nil
}

//shuffle card
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) 
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)		
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}