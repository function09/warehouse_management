package shipment

import (
	"database/sql"
	"fmt"
	"strconv"
)

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
		return nil, fmt.Errorf("Error querying shipments: %v", err)
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
		return nil, fmt.Errorf("No rows found: %v", err)
	}

	return shipments, nil
}

func (r *PostgreSQLRepository) UpdateShipments(dd string, rq int, id int) (string, error) {
	sqlStatement := "UPDATE shipments SET date_delivered=$1, received_pallet_qty=$2 WHERE id=$3"

	result, err := r.db.Exec(sqlStatement, dd, rq, id)

	if err != nil {
		return "", fmt.Errorf("Error updating shipments %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "", fmt.Errorf("error checking affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("No shipment found with ID %d to update", id)
	}

	return fmt.Sprintf("Succesfully updated shipment ID: " + strconv.Itoa(id)), nil
}
