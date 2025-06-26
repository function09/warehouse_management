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
	Message string     `json:"message"`
	Code    int        `json:"code"`
	Data    []*Product `json:"data"`
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

	var pList []*Product

	p, err := h.productService.GetProductByID(IDInt)

	if err != nil {
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if p == nil {
		sendJSONError(w, "Product not found", http.StatusNotFound)
		return
	}

	pList = append(pList, p)

	data := SuccessMessage{
		Message: "Successfully fetched product with ID",
		Code:    http.StatusOK,
		Data:    pList,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode json: %v", err)
	}
}

func (h *ProductHandler) GetProductByName(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")

	if name == "" {
		sendJSONError(w, "Name cannot be blank", http.StatusBadRequest)
		return
	}

	// var pList []*Product

	p, err := h.productService.GetProductByName(name)

	if err != nil {
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// pList = append(pList, p)

	data := SuccessMessage{
		Message: "Successfully fetched product by name",
		Code:    http.StatusOK,
		Data:    p,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode json: %v", err)
	}

}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	limit := params.Get("limit")
	offset := params.Get("offset")

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		log.Printf("Error converting string to int: %v", err)
	}

	offsetInt, err := strconv.Atoi(offset)

	if err != nil {
		log.Printf("Error converting string to int: %v", err)
	}

	if limit == "" {
		sendJSONError(w, "No limit value specified", http.StatusBadRequest)
		return
	}

	if offset == "" {
		sendJSONError(w, "No offset value specified", http.StatusBadRequest)
		return
	}

	p, err := h.productService.GetAllProducts(limitInt, offsetInt)

	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := SuccessMessage{
		Message: "Successfully fetched product with Name",
		Code:    http.StatusOK,
		Data:    p,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode json: %v", err)
	}

}
