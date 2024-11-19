package main

func main() {
	// := は新変数を定義(初期化)するときにのみ使用. func 内で使用すること.
	// 既存の変数に新しい値を代入する場合は、= でよい
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
