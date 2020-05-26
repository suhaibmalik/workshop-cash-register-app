package cart

import (
	"net/http"

	"github.com/suhaibmalik/workshop-cash-register-app/pkg/store"
)

func HandleListItems(s *store.Store) http.HandlerFunc {
	return func(w *http.ResponseWriter, r http.Request) {
		// I have access to the store here with s!
	}
}
