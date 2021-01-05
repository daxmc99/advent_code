package main

import "testing"

func TestFindRow(t *testing.T) {
	got := findRow("FBFBBFF")
	if got != 44 {
		t.Errorf("expected 44, got %d", got)
	}
}
