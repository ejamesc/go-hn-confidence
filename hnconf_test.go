package main_test

import (
	"fmt"
	"math"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/ejamesc/go-hn-confidence"
)

// Test Load function
func TestScrape(t *testing.T) {
	newsItems, err := main.Scrape()
	ok(t, err)

	for i, item := range newsItems {
		assert(t, item.Title != "", "no title detected for item %v", i)
		assert(t, item.Link != "", "no link detected for item %v", i)
		assert(t, item.Points >= 0, "invalid points number for item %v", i)
		assert(t, item.Comments >= 0, "invalid comments number for item %v", i)
		if i < len(newsItems)-1 {
			assert(t, item.Score() >= newsItems[i+1].Score(), "result is not sorted")
		}
	}
}

// Test Wilson Score calculation
func TestScore(t *testing.T) {
	mockItem := &main.NewsItem{Points: 100, Comments: 23}
	equals(t, 0.735, round(mockItem.Score(), 3))

	mockItem2 := &main.NewsItem{Points: 123, Comments: 123}
	equals(t, 0.438, round(mockItem2.Score(), 3))
}

// HELPERS
func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	up := math.Floor((f * shift) + .5)
	return up / shift
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
