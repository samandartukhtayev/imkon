package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{db: db}
}

func (ctgr *categoryRepo) CreateCategory(req *repo.CreateCategoryReq) (*repo.Category, error) {
	var result repo.Category
	query := `
		INSERT into categories(
			name
		)VALUES($1)
		RETURNING
			id,
			name
	`
	row := ctgr.db.QueryRow(query, req.Name)
	if err := row.Scan(
		&result.Id,
		&result.Name,
	); err != nil {
		return &repo.Category{}, err
	}
	return &result, nil
}

func (ctgr *categoryRepo) GetAllCategories(req *repo.GetAllCategoriesReq) (*repo.GetAllCategoriesRes, error) {
	var result repo.GetAllCategoriesRes
	// TODO:
	return &result, nil
}
