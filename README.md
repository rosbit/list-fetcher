# list-fetcher, a utility to fetch list page by page and change the results to an iterator

## Interface to implement

```go
import (
	"encoding/json"
)

type PageFetcher interface {
    GetNextPage() (total int64, list []json.RawMessage, err error)
    ErrorOccurrs(err error)
}
```

### Utility functions

```go
func FetchList(pf PageFetcher) (total int64, it <-chan json.RawMessage, err error) {
    //
}

func Dump(w io.Writer, it <-chan json.RawMessage) {
    //
}
func DumpJSON(w io.Writer, it <-chan json.RawMessage) {
    //
}
```

### Usage

See [list-fetcher_test.go](list-fetcher_test.go)
