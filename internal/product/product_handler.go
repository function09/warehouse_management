package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productService *Service
}

func NewProductHandler(ps *Service) *ProductHandler {
	return &ProductHandler{productService: ps}
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessMessage struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Data    *Product `json:"data"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := ErrorMessage{
		Message: message,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("failed to encode json: %v", err)
	}
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	ID := params.Get("id")
	IDInt, err := strconv.Atoi(ID)

	if err != nil {
		log.Printf("Error converting ID to int: %v", err)
		return
	}

	if ID == "" {
		sendJSONError(w, "ID cannot be blank", http.StatusBadRequest)
		return
	}

	if IDInt <= 0 {
		sendJSONError(w, "ID cannot be 0 or a negative value", http.StatusBadRequest)
		return
	}

	p, err := h.productService.GetProductByID(IDInt)

	if err != nil {
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
	}

	if p == nil {
		sendJSONError(w, "Product not found", http.StatusNotFound)
		return
	}

	data := SuccessMessage{
		Message: "Successfully fetched product with ID",
		Code:    http.StatusOK,
		Data:    p,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode json: %v", err)
	}
}
