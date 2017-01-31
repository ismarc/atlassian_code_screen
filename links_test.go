package main

import (
	"testing"
)

func TestLinksParsing(t *testing.T) {
	result := parseLinks("@chris you around?")
	if len(result) != 0 {
		t.Fail()
	}

	result = parseLinks("Good morning! (megusta) (coffee)")

	if len(result) != 0 {
		t.Fail()
	}

	result = parseLinks("Olympics are starting soon; http://www.nbcolympics.com")

	if len(result) != 1 {
		t.Fail()
	} else {
		if result[0].Url != "http://www.nbcolympics.com" {
			t.Fail()
		}
		if result[0].Title != "2018 PyeongChang Olympic Games | NBC Olympics" {
			t.Fail()
		}
	}

	result = parseLinks("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016")

	if len(result) != 1 {
		t.Fail()
	} else {
		if result[0].Url != "https://twitter.com/jdorfman/status/430511497475670016" {
			t.Fail()
		}

		if result[0].Title != "Justin Dorfman on Twitter: &quot;nice @littlebigdetail from @HipChat (shows hex colors when pasted in chat). http://t.co/7cI6Gjy5pq&quot;" {
			t.Fail()
		}
	}
}
