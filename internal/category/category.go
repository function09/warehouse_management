package category

type Category struct {
	CategoryID   int    `json:"ID"`
	CategoryName string `json:"category"`
}

type Repository interface {
	GetCategoryByName(n string) (*Category, error)
	GetCategoryByID(id int) (*Category, error)
	AddNewCategory(n string) (int64, error)
	UpdateCategory(n string, id int) (int64, error)
	DeleteCategory(id int) (int64, error)
}
