package main

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	errOsRemove := os.Remove("_decktesting")
	if errOsRemove != nil {
		fmt.Println("error os remove", errOsRemove)
	}
	deck := newDeck()
	errSaveToFile := deck.saveToFile("_decktesting")
	if errSaveToFile != nil {
		fmt.Println("error read from file", errSaveToFile)
	}

	loadedDeck, errFromFile := newDeckFromFile("_decktesting")
	if errFromFile != nil {
		fmt.Println("error read from file", errFromFile)
	}

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	errosRemove := os.Remove("_decktesting")
	if errosRemove != nil {
		fmt.Println("error os remove", errosRemove)
	}
}

// testing use convey
func TestDummy(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 16
		Convey("When the integer is incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 17)
			})
		})

	})
}

//done
func TestNewDeckCard(t *testing.T) {
	d := newDeck()
	Convey("Given a fresh deck", t, func() {
		Convey("Expected deck length of 16", func() {
			So(len(d), ShouldEqual, 16)
		})
		Convey("Expected first card Ace of Spades", func() {
			So(d[0], ShouldEqual, "Ace of Spades")
		})
		Convey("Expected last card Four of Clubs", func() {
			So(d[len(d)-1], ShouldEqual, "Four of Clubs")
		})
	})
}

//done
func TestNewDeckPrint(t *testing.T) {
	d := newDeck()
	d.print()
	Convey("Given all deck print", t, func() {
		Convey("Expected deck length of 16", func() {
			So(len(d), ShouldEqual, 16)
		})
	})
}

//done
func TestDealCard(t *testing.T) {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	Convey("Given deal card", t, func() {

		Convey("Expected hand cards length of 5", func() {
			So(len(hand), ShouldEqual, 5)
		})

		Convey("Expected hand first card Ace of Spades", func() {
			So(hand[0], ShouldEqual, "Ace of Spades")
		})

		Convey("Expected hand last card Ace of Diamonds", func() {
			So(hand[len(hand)-1], ShouldEqual, "Ace of Diamonds")
		})

		Convey("Expected remaining cards length of 11", func() {
			So(len(remainingCards), ShouldEqual, 11)
		})

		Convey("Expected remaining first card Two of Diamonds", func() {
			So(remainingCards[0], ShouldEqual, "Two of Diamonds")
		})

		Convey("Expected remaining last card Four of Clubs", func() {
			So(remainingCards[len(remainingCards)-1], ShouldEqual, "Four of Clubs")
		})

	})
}

//done
func TestNewDeckShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	foundItem := false
	var keyWord = "Ace of Spades"
	for i := range d {
		if d[i] == keyWord {
			foundItem = true
		}
	}

	Convey("Expected shuffle card", t, func() {
		Convey("Expected deck length of 16", func() {
			So(len(d), ShouldEqual, 16)
		})
		Convey("Expected found card Ace of Spades ", func() {
			So(foundItem, ShouldBeTrue)
		})
	})
}

//done
func TestSaveToDeckNewDeckFromFile(t *testing.T) {
	errOsRemove := os.Remove("_decktesting")
	if errOsRemove != nil {
		fmt.Println("error os remove", errOsRemove)
	}
	deck := newDeck()
	errSaveToFile := deck.saveToFile("_decktesting")
	if errSaveToFile != nil {
		fmt.Println("error read from file", errSaveToFile)
	}

	loadedDeck, errFromFile := newDeckFromFile("_decktesting")
	if errFromFile != nil {
		fmt.Println("error read from file", errFromFile)
	}

	Convey("Testing save to file", t, func() {
		Convey("Expected deck length of 16", func() {
			So(len(loadedDeck), ShouldEqual, 16)
		})
		Convey("Expected no error", func() {
			So(errFromFile, ShouldEqual, nil)
		})
	})

	errosRemove := os.Remove("_decktesting")
	if errosRemove != nil {
		fmt.Println("error os remove", errosRemove)
	}

}
