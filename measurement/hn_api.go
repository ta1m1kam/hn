package main

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/opengraph"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HackerNews struct {
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string
}

func GetHackerNewsDetail(ids []int) ([]HackerNews, error) {
	var hns []HackerNews
	var hn HackerNews
	for _, s := range ids {
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		res, err := http.Get(url)
		if err != nil {
			continue
		}

		err = json.NewDecoder(res.Body).Decode(&hn)
		if err != nil {
			continue
		}

		if hn.URL != "" {
			og, err := opengraph.Fetch(hn.URL)
			if err != nil {
				continue
			}
			hn.Description = og.Description
		}

		hns = append(hns, hn)
		res.Body.Close()
	}

	return hns, nil
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

	n := 20
	start := time.Now()
	GetHackerNewsDetail(idHn[0 : n-1])
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
