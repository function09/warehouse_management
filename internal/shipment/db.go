package shipment

import "database/sql"

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgresSQLRepository(db *sql.DB) Repository {
	return &PostgreSQLRepository{db: db}
}

func (r *PostgreSQLRepository) GetShipments(l int, o int) ([]*Shipment, error) {
	sqlStatement := "SELECT id, po_number, status, expected_delivery, date_delivered, pallet_qty, received_pallet_qty, supplier_id FROM shipments ORDER BY id LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(sqlStatement, l, o)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var shipments []*Shipment

	for rows.Next() {
		var shipment Shipment

		if err := rows.Scan(&shipment.ID, &shipment.PO_Number, &shipment.Status, &shipment.Expected_delivery, &shipment.Date_delivered, &shipment.Pallet_qty, &shipment.Received_pallet_qty, &shipment.Supplier_id); err != nil {
			return nil, err
		}

		shipments = append(shipments, &shipment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shipments, nil
}
