package store

import (
	"encoding/json"

	"github.com/suhaibmalik/workshop-cash-register-app/pkg/item"
)

type Store struct {
	Items []item.Item `json:"items"`
}

type Totals struct {
	// todo
}

func (s *Store) MarshalJSON() ([]byte, error) {
	t := Totals{} // todo:

	return json.Marshal(&struct {
		Items  []item.Item `json:"items"`
		Totals Totals
	}{
		Items:  s,
		Totals: t,
	})
}
