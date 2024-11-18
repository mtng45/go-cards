package main

import "fmt"

func main() {
	// var card string = "Ace of Spades". func 内でなくてもOK.
	// := は新変数を定義(初期化)するときにのみ使用. func 内で使用すること.
	// 既存の変数に新しい値を代入する場合は、= でよい
	card := newCard()
	fmt.Println(card)

	cards := deck{"Ace of Diamonds", newCard()}
	// append は Slice を変更せず、新しい Slice を返す => 重要
	cards = append(cards, "Six of Spades")

	cards.print() // deck 型に print メソッドを定義しているため、これが使用できる
}

func newCard() string {
	return "Five of Diamonds"
}
