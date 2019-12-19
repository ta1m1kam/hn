package main

import (
	"encoding/json"
	"github.com/k0kubun/pp"
	"github.com/otiai10/opengraph"

	//"github.com/otiai10/opengraph"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type HackerNews struct {
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string
}

func GetHackerNewsDetail(ids []int) []HackerNews {
	wg := new(sync.WaitGroup)
	var hns []HackerNews
	var chn = make(chan HackerNews)
	for _, s := range ids {
		wg.Add(1)
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		go func(url string) {
			var hn HackerNews
			res, _ := http.Get(url)
			body, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(body, &hn)
			if hn.Url != "" {
				og, err := opengraph.Fetch(hn.Url)
				if err != nil {
					log.Fatal(err)
				}
				hn.Description = og.Description
			}
			chn <- hn
			wg.Done()
		}(url)
		hns = append(hns, <-chn)
	}
	pp.Print(hns)
	wg.Wait()
	return hns
}

func GetHackerNews(n int) []HackerNews {
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

	//var hns []HackerNews
	hns := GetHackerNewsDetail(idHn[0 : n-1])
	return hns
}
