package repo

type Category struct {
	Id   int64
	Name string
}

type CreateCategoryReq struct {
	Name string
}

type GetAllCategoriesReq struct {
	Page       int64
	Limit      string
	SortByName string
}
type GetAllCategoriesRes struct {
	Categories []*Category
	Count      int64
}

type CategoryStorageI interface {
	CreateCategory(*CreateCategoryReq) (*Category, error)
	GetAllCategories(*GetAllCategoriesReq) (*GetAllCategoriesRes, error)
}
