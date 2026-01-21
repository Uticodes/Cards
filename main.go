package main

import "fmt"

// func main() {
// 	// var card string = "Ace of Spades"
// 	card := newCard()
// 	fmt.Println(card)
// }

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func main() {
// 	cards := deck{"Ace of Diamonds", getTitle(), "Six of Spades"}
// 	cards = append(cards, "Seven of Hearts")
// 	cards.print()
// }

//	func getTitle() string {
//		return "Harry Potter"
//	}
type book string

func (b book) printTitle() {
	fmt.Println(b)
}

func main() {
	// var b book = "Harry Potter"
	// b.printTitle()
	// cards := newDeck()
	// cards.saveToFileOs("my_os_cards_ss")
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// err := cards.saveToFile("my_cards")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// fmt.Println("-----")
	// fmt.Println([]byte(cards.toString()))
	// cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println("----")
	// remainingDeck.print()

	// color1, color2, color3 := colors()

	// fmt.Println(color1, color2, color3)
	// greeting := "Hello, World!"
	// fmt.Println([]byte(greeting))
}
