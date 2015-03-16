package main

import (
	"testing"
)

func TestDecodeText(t *testing.T) {
	wordData := []byte{0x83, 0x65, 0x83, 0x58, 0x83, 0x67}
	word := string(wordData[:])
	if decodeText(word, "shiftjis") != "テスト" {
		t.Fail()
	}
}
