package store

import (
	"encoding/json"
	"sync"

	"github.com/suhaibmalik/workshop-cash-register-app/pkg/item"
)

type Store struct {
	mu    sync.Mutex
	Items []item.Item
}

func (s *Store) AddItem(item item.Item) {
	s.mu.Lock()
	// todo: add item
	s.mu.Unlock()
}

type Totals struct {
	// todo
}

func (s *Store) MarshalJSON() ([]byte, error) {
	t := Totals{} // todo: calculate totals

	return json.Marshal(&struct {
		Items  []item.Item `json:"items"`
		Totals Totals      `json:"totals"`
	}{
		Items:  s.Items,
		Totals: t,
	})
}
