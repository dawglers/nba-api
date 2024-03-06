package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ndesai96/nba-api/http"
)

const NBAStatsBaseURL = "https://stats.nba.com/stats/"

type Base struct {
	params  url.Values
	results *Results
}

func (b *Base) Request(path Path) error {
	// Build an HTTP url
	requestURL := fmt.Sprint(NBAStatsBaseURL, path)

	url, err := url.Parse(requestURL)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}

	url.RawQuery = b.params.Encode()

	response, err := http.MakeRequest(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	var r Results

	// Unmarshal JSON into the map
	err = json.Unmarshal(body, &r)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	b.results = &r

	return nil
}
