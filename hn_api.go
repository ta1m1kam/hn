package main

import (
	"encoding/json"
	"github.com/otiai10/opengraph"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type HackerNews struct {
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string
}

func GetHackerNews() []HackerNews {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var idHn []int
	json.Unmarshal(body, &idHn)

	var hns []HackerNews
	var hn HackerNews
	cnt := 0
	for _, s := range idHn {
		if cnt > 29 {
			break
		}
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		res, _ := http.Get(url)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &hn)
		og, err := opengraph.Fetch(hn.Url)
		if err != nil {
			log.Fatal(err)
		}
		hn.Description = og.Description
		hns = append(hns, hn)
		cnt += 1
	}

	return hns
}
