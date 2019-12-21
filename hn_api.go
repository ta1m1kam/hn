package main

import (
	"encoding/json"

	"github.com/otiai10/opengraph"

	//"github.com/otiai10/opengraph"
	"io/ioutil"
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

func GetHackerNewsDetail(ids []int) ([]HackerNews, error) {
	wg := new(sync.WaitGroup)
	var hns []HackerNews
	var chn = make(chan HackerNews)
	for _, s := range ids {
		wg.Add(1)
		url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"
		go func(url string) {
			var hn HackerNews
			res, err := http.Get(url)
			if err != nil {
				return
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return
			}
			err = json.Unmarshal(body, &hn)
			if err != nil {
				return
			}
			if hn.Url != "" {
				og, err := opengraph.Fetch(hn.Url)
				if err != nil {
					return
				}
				hn.Description = og.Description
			}
			chn <- hn
			wg.Done()
		}(url)
		hns = append(hns, <-chn)
	}
	defer close(chn)
	wg.Wait()
	return hns, nil
}

func GetHackerNews(n int) ([]HackerNews, error) {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var idHn []int
	err = json.Unmarshal(body, &idHn)
	if err != nil {
		return nil, err
	}

	//var hns []HackerNews
	return GetHackerNewsDetail(idHn[0 : n-1])
}
