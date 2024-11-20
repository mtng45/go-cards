package deck

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append は Slice を変更せず、新しい Slice を返す
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	// プロジェクトルートディレクトリのパスを取得
	rootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error: failed to get root directory: %w", err)
	}

	// out ディレクトリが存在することを確認する
	if err := os.MkdirAll(filepath.Join(rootDir, "out"), 0755); err != nil {
		return fmt.Errorf("Error:failed to create out directory: %w", err)
	}

	// ファイルパスを作成する
	filepath := filepath.Join("out", filename)

	return os.WriteFile(filepath, []byte(d.toString()), 0666) // 0666 はパーミッション
}

func NewDeckFromFile(filename string) deck {
	filepath := filepath.Join("out", filename)
	bs, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
