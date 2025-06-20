package api

import (
	"net/http"

	service "github.com/function09/warehouse_management/service/products"
)

type ProductHandler struct {
	productService *service.Service
}

func NewProductHandler(ps *service.Service) *ProductHandler {
	return &ProductHandler{productService: ps}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	ID := params.Get("id")

	if ID == "" {

	}
}
