package main

import "testing"

func TestFindRow(t *testing.T) {
	got := findRow("FBFBBFF")
	if got != 44 {
		t.Errorf("expected 44, got %d", got)
	}
}

func TestFindCol(t *testing.T) {
	got := findCol("RLR")
	if got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}
