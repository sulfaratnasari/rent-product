package lostanimal

import (
	"encoding/json"
	"fmt"
	"net/http"
	repouc "rent-product/internal/interface/usecase"
	"strconv"
)

type ProductHandler struct {
	ProductUC repouc.ProductUC
}

func New(handler *ProductHandler) *ProductHandler {
	return handler
}

func (h *ProductHandler) ProductAvailability(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	productID, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	//productID, _ = strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	productAvailability, err := h.ProductUC.ProductAvailabilityList(ctx, productID, month, year)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	responseJson, err := json.Marshal(productAvailability)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
