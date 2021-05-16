package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Card struct {
	number int // The number of the card in sequence (A=1, J=11, Q=12, etc.)
	suit   string
	symbol string // A, 2, 3, ... , J, Q, K
}

// Create a new factory default deck of cards, no jokers.
func NewDeck() [52]Card {
	var newDeck [52]Card

	for s := 0; s <= 12; s++ {
		newDeck[s].number = s + 1
		newDeck[s].suit = "Spades"
	}
	for d := 13; d <= 26; d++ {
		newDeck[d].number = (d + 1) - 13
		newDeck[d].suit = "Diamonds"
	}
	for c := 26; c <= 39; c++ {
		newDeck[c].number = (c + 1) - 26
		newDeck[c].suit = "Clubs"
	}
	for h := 39; h <= 51; h++ {
		newDeck[h].number = (h + 1) - 39
		newDeck[h].suit = "Hearts"
	}
	fmt.Println("Created a new 52 card deck." + "\n")
	time.Sleep(2 * time.Second)
	FormatDeck(&newDeck)
	return newDeck
}

func FormatDeck(deck *[52]Card) { // Delegate symbols
	for i := 0; i < len(deck); i++ {
		if (deck[i].number != 1) && (deck[i].number != 11) &&
			(deck[i].number != 12) && (deck[i].number != 13) {
			deck[i].symbol = strconv.Itoa(deck[i].number)
		} else {
			switch deck[i].number {
			case 1:
				deck[i].symbol = "A"
			case 11:
				deck[i].symbol = "J"
			case 12:
				deck[i].symbol = "Q"
			case 13:
				deck[i].symbol = "K"
			default:
				deck[i].symbol = strconv.Itoa(deck[i].number)
			}
		}
	}
}

func PrintDeck(deck [52]Card) {
	cardMap := make(map[int]string)
	cardMap[1] = "A"
	cardMap[11] = "J"
	cardMap[12] = "Q"
	cardMap[13] = "K"
	suitMap := make(map[string]rune) // Unicode suit characters
	suitMap["Spades"] = '\u2660'
	suitMap["Diamonds"] = '\u2666'
	suitMap["Clubs"] = '\u2663'
	suitMap["Hearts"] = '\u2665'

	for i := 0; i < len(deck); i++ { // Display the cards
		fmt.Printf("%v"+" "+"%q"+"\n",
			deck[i].symbol, suitMap[deck[i].suit])
	}
	fmt.Print("\n")
}

func ShuffleDeck(deck [52]Card) [52]Card {
	r := rand.New(rand.NewSource(time.Now().Unix())) // Seed
	shuffled := deck
	for i, randIndex := range r.Perm(len(deck)) {
		shuffled[i] = deck[randIndex]
	}
	fmt.Println("The deck has been shuffled." + "\n")
	time.Sleep(2 * time.Second)
	return shuffled
}

func main() {
	deck := NewDeck() // initial deck

	for {
		var input string

		fmt.Println("Enter one of the following keys: " + "\n")
		fmt.Println("D ................ Display deck" + "\n" +
			"S ................ Shuffle deck" + "\n" +
			"N ................ New deck" + "\n" +
			"Q ................ Quit" + "\n")

		fmt.Scanln(&input)

		switch input {
		case "D":
			PrintDeck(deck)
		case "S":
			deck = ShuffleDeck(deck)
		case "N":
			deck = NewDeck()
		case "Q":
			os.Exit(0)
		default:
			fmt.Println("\n" + "Unknown command. Please try again." + "\n")
			input = "INVALID"
		}
		if input == "INVALID" {
			time.Sleep(2 * time.Second)
		}
	}
}
