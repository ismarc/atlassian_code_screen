package main

import (
	"testing"
)

func TestEmoticonParsing(t *testing.T) {
	result := parseEmoticons("@chris you around?")
	if len(result) > 0 {
		t.Fail()
	}

	result = parseEmoticons("Good morning! (megusta) (coffee)")
	if len(result) != 2 {
		t.Fail()
	} else {
		if result[0] != "megusta" {
			t.Fail()
		}

		if result[1] != "coffee" {
			t.Fail()
		}
	}

	result = parseEmoticons("Olympics are starting soon; http://www.nbcolympics.com")
	if len(result) > 0 {
		t.Fail()
	}

	result = parseEmoticons("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016")
	if len(result) != 1 {
		t.Fail()
	} else {
		if result[0] != "success" {
			t.Fail()
		}
	}
}
