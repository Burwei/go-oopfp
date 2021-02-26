package liboop

import (
	"reflect"
	"testing"
)

func TestBasePlayer_addResult(t *testing.T) {
	b := BasePlayer{}
	arg := [2]int8{1, 2}
	b.addResult(arg)
	if reflect.DeepEqual(b.results, []string{"1A2B"}) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}

func TestBasePlayer_getID(t *testing.T) {
	b := BasePlayer{}
	b.setID(12)
	id := b.getID()
	if id == b.id {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}

func TestBasePlayer_setID(t *testing.T) {
	b := BasePlayer{}
	b.setID(12)
	if b.id == 12 {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}

func TestBasePlayer_getGuesses(t *testing.T) {
	b := BasePlayer{}
	b.guesses = append(b.guesses,"1234")
	if reflect.DeepEqual(b.getGuesses(), []string{"1234"}) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}

func TestBasePlayer_getResults(t *testing.T) {
	b := BasePlayer{}
	arg := [2]int8{1, 2}
	b.addResult(arg)
	if reflect.DeepEqual(b.getResults(), []string{"1A2B"}) {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}
