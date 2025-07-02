package product

type Product struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}

type Repository interface {
	GetProductByID(id int) (*Product, error)
	GetProductByName(n string) ([]*Product, error)
	GetProductsByCategory(n string) ([]*Product, error)
	GetAllProducts(limit int, offset int) ([]*Product, error)
	AddNewProduct(name string, stock int) (int64, error)
}
