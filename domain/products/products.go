package products

type Product struct {
	ID       int
	Title    string
	Category string
	Stock    int
}

type Repository interface {
	GetByID(id int) (*Product, error)
}
