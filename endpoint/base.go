package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ndesai96/nba-api/http"
)

const NBAStatsBaseURL = "https://stats.nba.com/stats/"

type Endpoint interface {
	ParamSetter

	Request() error
	GetResults() *Results
}

type base struct {
	params  url.Values
	results *Results
	path    Path
}

func New(path Path) *base {
	return &base{
		params: url.Values{},
		path:   path,
	}
}

func (b *base) SetPath(path Path) {
	b.path = path
}

func (b *base) Request() error {
	if b.path == "" {
		return fmt.Errorf("error: path must be defined")
	}

	// Build an HTTP url
	requestURL := fmt.Sprint(NBAStatsBaseURL, b.path)

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

func (b *base) GetResults() *Results {
	return b.results
}
