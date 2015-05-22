package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ROOT_URL string = "https://news.ycombinator.com"

type NewsItem struct {
	Title    string
	Link     string
	Points   int
	Comments int
}

func (ni *NewsItem) Score() float64 {
	n := float64(ni.Points + ni.Comments)
	pos := float64(ni.Points)
	if n == 0 {
		return 0
	}

	z := 1.96 // we assume a 95% confidence interval
	phat := 1.0 * pos / n
	res := (phat + z*z/(2*n) - z*math.Sqrt((phat*(1-phat)+z*z/(4*n))/n)) / (1 + z*z/n)
	return res
}

type Items []*NewsItem

func (m Items) Len() int {
	return len(m)
}

func (m Items) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Items) Less(i, j int) bool {
	return m[i].Score() < m[j].Score()
}

func Load() ([]*NewsItem, error) {
	doc, err := goquery.NewDocument(ROOT_URL)
	if err != nil {
		return nil, err
	}

	res := []*NewsItem{}

	doc.Find(".athing").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title a").Text()
		link, exists := s.Find(".title a").Attr("href")
		if !exists {
			link = ""
		}

		if strings.HasPrefix(link, "item?") {
			link = ROOT_URL + "/" + link
		}
		item := &NewsItem{Title: title, Link: link}
		res = append(res, item)
		//fmt.Printf("%v - %v\n", title, link)
	})

	doc.Find(".subtext").Each(func(i int, s *goquery.Selection) {
		pString := s.Find(".score").Text()
		cString := s.Find("a").Last().Text()
		points := 0
		comments := 0

		if pString != "" {
			pSt := strings.Split(pString, " ")[0]
			points, err = strconv.Atoi(pSt)
			if err != nil {
				points = 0
			}
		}

		if cString != "" && cString != "discuss" {
			cSt := strings.Split(cString, " ")[0]
			comments, err = strconv.Atoi(cSt)
			if err != nil {
				comments = 0
			}
		}

		item := res[i]
		item.Points = points
		item.Comments = comments
	})

	sort.Sort(sort.Reverse(Items(res)))
	return res, nil
}

func main() {
	newsItems, err := Load()
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range newsItems {
		fmt.Printf("[%v] %v : %v - %v | %v\n", item.Score(), item.Points, item.Comments, item.Title, item.Link)
	}
}
