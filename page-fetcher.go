package lf

import (
	"encoding/json"
)

type PageFetcher interface {
	GetNextPage() (total int64, list []json.RawMessage, err error)
}
