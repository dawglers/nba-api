package endpoint

type Results struct {
	Resource   string         `json:"resource"`
	ResultSets []ResultSet    `json:"resultSets"`
	Parameters map[string]any `json:"parameters"`
}

type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}
