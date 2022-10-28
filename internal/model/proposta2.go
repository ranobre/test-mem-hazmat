package models

type Proposta2 struct {
	Attributes map[string][]interface{} `json:"attributes"`
}

type HazmatAttribute2 struct {
	Name       string      `json:"name"`
	Quantity   int         `json:"quantity"`
	Properties []TypeValue `json:"properties"`
}

type TypeValue struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}
