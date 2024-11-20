package main

import "cards/internal/deck"

func main() {
	// := は新変数を定義(初期化)するときにのみ使用. func 内で使用すること.
	// 既存の変数に新しい値を代入する場合は、= でよい
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
	cards.SaveToFile("my_cards")
}
