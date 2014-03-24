package main

import (
	"testing"
	_ "github.com/voxelbrain/osxcrossfix"
)

func TestRandom(t *testing.T) {
	if random() != 4 {
		t.Fatalf("This doesn't look random to me")
	}
}
