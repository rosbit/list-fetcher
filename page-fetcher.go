package lf

import (
	"encoding/json"
	"log"
)

type PageFetcher interface {
	GetNextPage() (total int64, list []json.RawMessage, err error)
	AdjustPage(list []json.RawMessage)
	ErrorOccurrs(err error)
}

type PageFetcherAdapter struct {
	Page int
	Offset int
}
func (a *PageFetcherAdapter) GetNextPage() (total int64, list []json.RawMessage, err error) {
	return
}
func (a *PageFetcherAdapter) AdjustPage(list []json.RawMessage) {
	a.Page += 1
	a.Offset += len(list)
}
func (a *PageFetcherAdapter) ErrorOccurrs(err error) {
	log.Printf("error occurs when calling GetNextPage() %d: %v\n", a.Page, err)
}
