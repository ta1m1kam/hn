package main

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/opengraph"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type HackerNews struct {
	n           int
	By          string `json:"by"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string
}

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

	//var hns []HackerNews
	n := 20
	start := time.Now()
	GetHackerNewsDetail(idHn[0 : n-1])
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
