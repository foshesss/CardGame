package cards

import (
	"testing"
)

func TestDerp(t *testing.T) {
	deck := newDeck()
	deck.print()
	deck.saveToFile("DerpAFlerp")
	deck.print()
}
