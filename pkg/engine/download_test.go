package engine

import (
/*
	"context"
	"fmt"
	"net/http"
	"testing"
*/
)

/*
func TestAcast(t *testing.T) {
	url := "https://sphinx.acast.com/p/open/s/6414da1f9a87fc0011e7a8ee/e/68f7fd3edbf5027e4936781a/media.mp3"

	i := 0
	for i < 100 {

		fmt.Println(url)

		cl := &http.Client{}

		req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			t.Fatalf("failed making request: %v", err)
		}

		req.Header.Set("User-Agent", "podfetch/1.0")

		resp, err := cl.Do(req)
		if err != nil {
			t.Fatalf("failed getting %s: %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Done!")
			break
		}

		fmt.Println(resp.StatusCode)

		url = resp.Header.Get("Location")
		i++
	}
}
*/
