package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "one of summer" {
		t.Errorf("Expected one of summer but got %v", d[0])
	}

	if d[len(d)-1] != "three of lei" {
		t.Errorf("Expected three of lei but got %v", d[len(d)-1])
	}
}
func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	savedDeck := newDeckFromFile("_decktesting")

	if len(savedDeck) != 9 {
		t.Errorf("Expected got 9 but got %v", len(savedDeck))
	}

	os.Remove("_decktesting")

}
