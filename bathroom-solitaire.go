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
	t_start := time.Now()
	n_trials := 1000000
	num_left := make([]int, n_trials)
	num_winners := 0

	// loop over all trials
	for i := 0; i < n_trials; i++ {

		// build base deck
		deck_0 := buildDeck()
		check := false

		// shuffle deck
		deck_down := shuffle(deck_0)

		// draw first four cards
		deck_up := make([]string, 0)
		deck_down, deck_up = drawCard(deck_down, deck_up, 4)

		// eliminate cards until deck_down is emtpy
		for len(deck_down) > 0 {
			if len(deck_up) > 3 {
				deck_up, check = checkCards(deck_up, check)
				if check == false && len(deck_down) > 0 {
					deck_down, deck_up = drawCard(deck_down, deck_up, 1)
				}
			} else {
				for len(deck_up) < 4  && len(deck_down) > 0 {
					deck_down, deck_up = drawCard(deck_down, deck_up, 1)
				}
			}
		}

		// eliminate cards from deck_up until game is over
		if len(deck_up) > 3 {
			check = true
			for check == true {
				if len(deck_up) > 3 {
					deck_up, check = checkCards(deck_up, check)
				} else {
					check = false
				}
			}
		}

		// update num_left and num_winners
		num_left[i] = len(deck_up)
		if len(deck_up) == 0 {
			num_winners += 1
		}
	}

	// print summary
	elapsed := time.Since(t_start)
	fmt.Println()
	fmt.Printf("Number of trials: %d\n", n_trials)
	fmt.Printf("Number of winners: %d\n", num_winners)
	fmt.Printf("Win percentage: %.2f%%\n", 100.*float64(num_winners)/float64(n_trials))
	fmt.Printf("Runtime: %s\n\n", elapsed)
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

//==============================================================================
// Draw function
//==============================================================================
func drawCard(deck_down []string, deck_up []string, n int) ([]string, []string) {
	for i := 0; i < n; i++ {
		deck_up = append(deck_up, deck_down[:1]...)
		deck_down = append(deck_down[1:])
	}
	return deck_down, deck_up
}

//==============================================================================
// Check top 4 cards
//==============================================================================
func checkCards(deck_up []string, check bool) ([]string, bool) {
	test_cards := deck_up[len(deck_up)-4:]

	// check for same suit
	if string([]rune(test_cards[0])[1]) == string([]rune(test_cards[3])[1]) {
		deck_up = append(deck_up[:len(deck_up)-3], deck_up[len(deck_up)-1:]...)
		check = true
	// check for same card
	} else if string([]rune(test_cards[0])[0]) == string([]rune(test_cards[3])[0]) {
		deck_up = append(deck_up[:len(deck_up)-4])
		check = true
    } else {
		check = false
	}
	return deck_up, check
}
