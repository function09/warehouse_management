package category

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	categoryService *Service
}

func NewCategoryHandler(s *Service) *Handler {
	return &Handler{categoryService: s}
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessMessage struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []*Category `json:"data"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := ErrorMessage{
		Message: message,
		Code:    statusCode,
	}

	// Create a function to handle this altogether
	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("failed to encode json: %v", err)
	}
}
func (h *Handler) GetCategoryByName(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	category := params.Get("category")

	if category == "" {
		sendJSONError(w, "Category cannot be blank", http.StatusBadRequest)
	}

	var result []*Category

	cat, err := h.categoryService.GetCategoryByName(category)

	if err != nil {
		log.Printf("GetCategoryByName error: %v", err)
	}

	result = append(result, cat)

	data := SuccessMessage{
		Message: "Successfully fetched category",
		Code:    http.StatusOK,
		Data:    result,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
