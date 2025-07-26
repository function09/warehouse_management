package shipment

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ShipmentHandler struct {
	ShipmentService *Service
}

type SuccessMessage struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    []*Shipment `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func NewShipmentHandler(s *Service) *ShipmentHandler {
	return &ShipmentHandler{ShipmentService: s}
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

func (s *ShipmentHandler) GetShipments(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	limit := params.Get("limit")
	offset := params.Get("offset")

	if limit == "" {
		sendJSONError(w, "No limit value specified", http.StatusBadRequest)
		return
	}

	if offset == "" {
		sendJSONError(w, "No offset value specified", http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		sendJSONError(w, "Invalid 'limit' parameter", http.StatusBadRequest)
		return
	}

	offsetInt, err := strconv.Atoi(offset)

	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		sendJSONError(w, "Invalid 'offset' parameter", http.StatusBadRequest)
		return
	}

	shipments, err := s.ShipmentService.GetShipments(limitInt, offsetInt)

	if err != nil {
		log.Printf("GetShipments error: %v", err)
		sendJSONError(w, "Internal server error occured", http.StatusInternalServerError)
		return
	}

	message := SuccessMessage{
		Message: "Successfully retrieved all shipments",
		Status:  http.StatusOK,
		Data:    shipments,
	}

	w.Header().Set("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Printf("Failed to encode json: %v", err)
	}
}
