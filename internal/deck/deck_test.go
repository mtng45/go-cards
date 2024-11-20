package deck

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		// %v:どんな型の値でも適切に文字列化して表示できる
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	testPath := filepath.Join("out", testFile)

	// クリーンアップ
	os.Remove(testPath)

	deck := NewDeck()
	deck.SaveToFile(testFile)

	loadedDeck := NewDeckFromFile(testFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// クリーンアップ
	os.Remove(testPath)
}
