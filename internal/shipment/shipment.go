package shipment

import "time"

type Shipment struct {
	ID                  int        `json:"id"`
	PO_Number           string     `json:"po_number"`
	Status              string     `json:"status"`
	Expected_delivery   string     `json:"expected_delivery"`
	Date_delivered      *time.Time `json:"date_delivered"`
	Pallet_qty          int        `json:"pallet_qty"`
	Received_pallet_qty *int       `json:"received_pallet_qty"`
	Supplier_id         string     `json:"supplier_id"`
}

type Repository interface {
	GetShipments(l int, o int) ([]*Shipment, error)
	UpdateShipments(dd string, rq int, id int) (string, error)
}
