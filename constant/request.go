package constant

type Request struct {
	Examples []Example `json:"examples"`
	Rules    []Rule    `json:"rules"`
	Input    string    `json:"input"`
}
