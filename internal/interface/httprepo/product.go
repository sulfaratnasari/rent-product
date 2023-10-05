package httprepo

import "net/http"

type ProductHandler interface {
	ProductAvailability(w http.ResponseWriter, r *http.Request)
}
