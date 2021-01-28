package main

import "testing"

func TestRepeated(t *testing.T) {
	g := group{numPeople: 1, uniqueCount: 3, duplicateCount: 3,
		people: []person{person{"abc"}}}
	got := allYes(g)
	if got != 3 {
		t.Error("wrong answer")
	}
	g = group{numPeople: 3, uniqueCount: 3, duplicateCount: 0,
		people: []person{person{"a"}, person{"b"}, person{"c"}}}
	got = allYes(g)
	if got != 0 {
		t.Error("wrong answer")
	}
	g = group{numPeople: 2, uniqueCount: -1, duplicateCount: -1,
		people: []person{person{"ab"}, person{"ac"}}}
	got = allYes(g)
	if got != 1 {
		t.Error("wrong answer")
	}
	g = group{numPeople: 4, uniqueCount: -1, duplicateCount: -1,
		people: []person{person{"a"}, person{"a"}, {"a"}, {"a"}}}
	got = allYes(g)
	if got != 1 {
		t.Error("wrong answer")
	}

}
