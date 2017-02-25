/*
bathroom_solitaire.go

Bathroom solitaire simulation
Game description: http://en.wikipedia.org/wiki/One-Handed_Solitaire

Tom Pavlak
2/24/17
*/

//==============================================================================
// Imports
//==============================================================================
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//==============================================================================
// Main program
//==============================================================================
func main() {
	rand.Seed(time.Now().UnixNano())

	// build base deck
	deck_0 := buildDeck()

	// shuffle deck
	deck := shuffle(deck_0)

	fmt.Println(deck_0)
	fmt.Println()
	fmt.Println(deck)
}


//==============================================================================
// Deck building function
//==============================================================================
func buildDeck() []string {
	suits := []string{"h", "d", "c", "s"}
	cards := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J",
		              "Q", "K"}

	deck := make([]string, 52)
	counter := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			deck[counter] = cards[j] + suits[i]
			counter += 1
		}
	}
	return deck
}


//==============================================================================
// Shuffle function
//==============================================================================
func shuffle(deck_in []string) []string {
	deck_out := make([]string, len(deck_in))
	perm := rand.Perm(len(deck_in))
	for i, v := range perm {
		deck_out[v] = deck_in[i]
	}
	return deck_out
}
