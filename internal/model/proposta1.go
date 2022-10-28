package models

type Proposta1 struct {
	Attributes map[string][]interface{} `json:"attributes"`
}

type HazmatAttribute struct {
	Name       string      `json:"name"`
	Quantity   int         `json:"quantity"`
	HazmatData interface{} `json:"hazmat_data"`
}

type BatteryData struct {
	BatteryType []string `json:"battery_type"`
	SaleFormat  []string `json:"sale_format"`
	SaleNumber  []string `json:"sale_number"`
}

type AerosolData struct {
	Volume      []string `json:"volume"`
	PackageType []string `json:"package_type"`
}
