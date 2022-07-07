package main

/*
[CardProject]
While this isn't actually original content-- I'm using an UDemy tutorial for this--
I am modifying it how I feel fit and using this as a learning opportunity.

Derp
*/

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()

	cards.saveToFile("Derp-Test")
}
