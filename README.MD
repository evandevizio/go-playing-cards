# go-playingcards
A playing card deck simulator in Go

## Summary
This is a program that simulates a deck of playing cards with functions to shuffle, reset, add decks, etc.
For classic card games such as Blackjack, for example, this could be adapted into an engine of sorts.

The deck is an array of "Card" structs, where each Card is comprised of a number, a symbol, and a suit.

Menu:
---
* **Display deck**: Prints the contents of the deck to the console
* **Shuffle deck**: Randomizes the order of the cards
* **New deck**: Resets the deck to a standard one of 52 cards
* **Add deck(s)**: Adds one or more decks to the current deck
