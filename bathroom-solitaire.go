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
)

//==============================================================================
// Main program
//==============================================================================
func main() {
	suits := []string{"h", "d", "c", "s"}
	cards := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	deck_0 := make([]string, 52)

	counter := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			deck_0[counter] = cards[j] + suits[i]
			counter += 1
		}
	}

  // shuffle deck
  deck := shuffle(deck_0, 7)

  fmt.Println(suits)
  fmt.Println(suits[:2])
  fmt.Println(cards)
  fmt.Println(deck_0)
  fmt.Println(deck)
  fmt.Println(shuffle(deck_0, 1))
}

//==============================================================================
// Function to shuffle deck
//==============================================================================
func shuffle(deck_in []string, n int) []string {
  deck_out := make([]string, len(deck_in))
  for i := 0; i < n; i++ {
      perm := rand.Perm(len(deck_in))
      for i, v := range perm {
          deck_out[v] = deck_in[i]
      }
      deck_in = deck_out
  }
  return deck_out
}
