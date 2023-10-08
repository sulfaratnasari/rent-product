package lostanimal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rent-product/internal/entity/product"
	repouc "rent-product/internal/interface/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductUC repouc.ProductUC
}

func New(handler *ProductHandler) *ProductHandler {
	return handler
}

// Get Productivity Data
func (h *ProductHandler) ProductAvailability(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	vars := mux.Vars(r)
	id := vars["id"]
	productID, _ := strconv.ParseInt(id, 10, 64)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

// Get Productivity Data
func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var product product.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ProductUC.AddProduct(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseJson, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (h *ProductHandler) ListProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("list product-----")
	products, err := h.ProductUC.ProductList(ctx)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	responseJson, err := json.Marshal(products)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
