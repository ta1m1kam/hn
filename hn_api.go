package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type HackerNews struct {
	By    string `json:"by"`
	Score int64  `json:"score"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

func main() {
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
	for _, s := range idHn {
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		res2, _ := http.Get(url)
		robots, _ := ioutil.ReadAll(res2.Body)
		json.Unmarshal(robots, &hn)
		hns = append(hns, hn)
	}
	fmt.Println(hns[0].Title)
}
