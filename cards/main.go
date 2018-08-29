package main

func main() {
	deck := readFromFile("my_cards2")
	deck.shuffle()
	deck.print()
	//fmt.Println(cards)
}
