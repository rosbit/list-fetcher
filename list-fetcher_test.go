package lf

import (
	"encoding/json"
	"testing"
	"fmt"
	"os"
)

func TestFetchList(t *testing.T) {
	total, it, err := FetchList(newItemPageFetcher(0, 2))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if total == 0 {
		fmt.Printf("total: %d\n", total)
	}
	Dump(os.Stdout, it)
}

type item struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var (
	items = []*item {
		&item{
			Name: "Alice",
			Age: 5,
		},
		&item{
			Name: "Bob",
			Age: 6,
		},
		&item{
			Name: "Cat",
			Age: 7,
		},
		&item{
			Name: "Dog",
			Age: 8,
		},
	}
)

// PageFetcher implementation
type ItemPageFetcher struct {
	*PageFetcherAdapter
	pageSize int
}
func newItemPageFetcher(page, pageSize int) *ItemPageFetcher {
	if page < 0 {
		page = 0
	}
	if pageSize <= 0 {
		pageSize = 3
	}
	return &ItemPageFetcher{
		PageFetcherAdapter: &PageFetcherAdapter{
			Page: page,
		},
		pageSize: pageSize,
	}
}

func (ipf *ItemPageFetcher) GetNextPage() (total int64, list []json.RawMessage, err error) {
	total = int64(len(items))

	start := ipf.Page * ipf.pageSize
	end := start + ipf.pageSize
	if start >= len(items) {
		err = fmt.Errorf("no more data")
		return
	}
	if end > len(items) {
		end = len(items)
	}
	jb := makeJSON(items[start:end])
	err = json.Unmarshal(jb, &list)
	return
}

func makeJSON(itemSlice []*item) ([]byte) {
	b, _ := json.Marshal(itemSlice)
	return b
}

