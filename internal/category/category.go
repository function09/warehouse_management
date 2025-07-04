package category

type Category struct {
	CategoryID   int    `json:"ID"`
	CategoryName string `json:"category"`
}

type Repository interface {
	GetCategoryByName(n string) (*Category, error)
	GetCategoryByID(id int) (*Category, error)
	AddCategory(n string) (int64, error)
	UpdateCategory(id int) (int64, error)
	DeleteCategory(id int) (int64, error)
}
