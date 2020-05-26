package item

type Item struct {
	Name string `json:"item"`
	UnitPrice float64 `json:"price_per_unit"`
	Units float64 `json:"units"`
}
