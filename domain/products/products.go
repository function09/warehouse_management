package products

type Product struct {
	ID       int
	Title    string
	Category string
	Stock    int
}

type Repository interface {
	GetProductByID(id int) (*Product, error)
	GetProductByName(n string) (*Product, error)
	GetAllProducts() ([]*Product, error)
}
