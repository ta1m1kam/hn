package main

import (
	"encoding/json"
	"sort"
	"sync"

	"github.com/otiai10/opengraph"

	"net/http"
	"strconv"
)

// HackerNews store entry on hackernews.
type HackerNews struct {
	n           int
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string
}

// GetHackerNewsDetail return entries's detail on hackernews.
func GetHackerNewsDetail(ids []int) ([]HackerNews, error) {
	var wg sync.WaitGroup
	var hns []HackerNews
	var chn = make(chan HackerNews, len(ids))

	for _, s := range ids {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()

			url := "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(s) + ".json?print=pretty"

			var hn HackerNews
			hn.n = s

			res, err := http.Get(url)
			if err != nil {
				return
			}
			defer res.Body.Close()

			err = json.NewDecoder(res.Body).Decode(&hn)
			if err != nil {
				return
			}
			if hn.URL != "" {
				og, err := opengraph.Fetch(hn.URL)
				if err != nil {
					return
				}
				hn.Description = og.Description
			}
			chn <- hn
		}(s)
	}
	wg.Wait()
	close(chn)

	for e := range chn {
		hns = append(hns, e)
	}

	sort.Slice(hns, func(i, j int) bool {
		return hns[i].n < hns[j].n
	})
	return hns, nil
}

// GetHackerNews return entries on hackernews.
func GetHackerNews(n int) ([]HackerNews, error) {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var idHn []int
	err = json.NewDecoder(res.Body).Decode(&idHn)
	if err != nil {
		return nil, err
	}

	//var hns []HackerNews
	return GetHackerNewsDetail(idHn[0 : n-1])
}
