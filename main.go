package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	var routes = Routes{
		Route{
			"Index",
			"POST",
			"/",
			Index,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

type Link struct {
	Url string `json:"url"`
	Title string `json:"title"`
}

type Message struct {
	Mentions []string `json:"mentions"`
	Emoticons []string `json:"emoticons"`
	Links []Link `json:"links"`
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	body := string(data)

	mentions := parseMentions(body)
	emoticons := parseEmoticons(body)
	links := parseLinks(body)
	message := Message {
		Mentions: mentions,
		Emoticons: emoticons,
		Links: links}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func parseMentions(data string) []string {
	// Cheat to handle mentions at end of message
	data = data + " "
	s := strings.Index(data, "@")

	if s > -1 {
		sub := data[s+1:]
		e := strings.Index(sub, " ")
		if e > -1 {
			token := sub[:e]
			result := []string{token}
			if e != len(sub) {
				return append(result, parseMentions(sub[e+1:])...)
			} else {
				return result
			}
		}
	}
	return []string{}
}

func parseEmoticons(data string) []string {
	s := strings.Index(data, "(")

	if s > -1 {
		sub := data[s+1:]
		e := strings.Index(sub, ")")
		if e > -1 {
			token := sub[:e]
			result := []string{token}
			if e != len(sub) {
				return append(result, parseEmoticons(sub[e+1:])...)
			} else {
				return result
			}
		}
	}
	return []string{}
}

func parseLinks(data string) []Link {
	// Cheat to handle links at end of message
	data = data + " "
	s := strings.Index(data, "http")

	if s > -1 {
		sub := data[s:]
		e := strings.Index(sub, " ")
		if e > -1 {
			token := sub[:e]
			result := []Link{
				Link{
					Url: token,
					Title: getPageTitle(token),
				},
			}
			if e != len(sub) {
				return append(result, parseLinks(sub[e+1:])...)
			} else {
				return result
			}
		}
	}

	return []Link{}
}

func getPageTitle(link string) string {
	resp, err := http.Get(link)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	data := string(body)

	s := strings.Index(data, "<title>")

	if s > -1 {
		sub := data[s+7:]
		e := strings.Index(sub, "</title>")
		if e > -1 {
			return sub[:e]
		}
	}
	return ""
}
