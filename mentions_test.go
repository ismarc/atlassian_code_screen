package main

import (
	"testing"
)

func TestMentionsParsing(t *testing.T) {
	result := parseMentions("@chris you around?")
	if len(result) != 1 {
		t.Fail()
	} else {
		if result[0] != "chris" {
			t.Fail()
		}
	}

	result = parseMentions("Good morning! (megusta) (coffee)")
	if len(result) > 0 {
		t.Fail()
	}

	result = parseMentions("Olympics are starting soon; http://www.nbcolympics.com")

	if len(result) > 0 {
		t.Fail()
	}

	result = parseMentions("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016")

	if len(result) != 2 {
		t.Fail()
	} else {
		if result[0] != "bob" {
			t.Fail()
		}
		if result[1] != "john" {
			t.Fail()
		}
	}
}
