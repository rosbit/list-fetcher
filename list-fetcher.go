package lf

import (
	"encoding/json"
)

func FetchList(pf PageFetcher) (total int64, it <-chan json.RawMessage, err error) {
	// get first page
	t, l, e := pf.GetNextPage()
	if e != nil {
		err = e
		return
	}
	if total = t; total <= 0 {
		return
	}

	listPage := make(chan []json.RawMessage)
	it = makeChanList(listPage)

	// get other pages
	go func() {
		count := int64(0)
		for {
			listPage <- l
			count += int64(len(l))
			if count >= total {
				break
			}

			if _, l, e = pf.GetNextPage(); e != nil {
				break
			}
		}

		close(listPage)
	}()

	return
}

func makeChanList(listPage <-chan []json.RawMessage) (<-chan json.RawMessage) {
	it := make(chan json.RawMessage)
	go func() {
		for l := range listPage {
			for i, _ := range l {
				it <- l[i]
			}
		}

		close(it)
	}()

	return it
}
