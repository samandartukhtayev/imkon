package models

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type GetAllCategoriesReq struct {
	Page       int64  `json:"page"  binding:"required" default:"1"`
	Limit      string `json:"limit" binding:"required" default:"10"`
	SortByName string `json:"sort_by_name"  binding:"required,oneof=superadmin user"`
}

type GetAllCategoriesRes struct {
	Categories []*Category `json:"categories"`
	Count      int64       `json:"count"`
}
