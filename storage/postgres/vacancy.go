package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/pkg/utils"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type vacancyRepo struct {
	db *sqlx.DB
}

func NewVacancy(db *sqlx.DB) repo.VacancyStorageI {
	return &vacancyRepo{db: db}
}

func ConvertVacancy(row *sql.Row) (*repo.VacancyResponse, error) {
	var result repo.VacancyResponse
	if err := row.Scan(
		&result.Id,
		&result.Name,
		&result.CategoryInfo.Id,
		&result.CategoryInfo.Name,
		&result.ImageUrl,
		&result.Address,
		&result.JobType,
		&result.Info,
		&result.MinSalary,
		&result.MaxSalary,
		&result.BusinessInfo.Id,
		&result.BusinessInfo.Name,
		&result.BusinessInfo.Address,
		&result.BusinessInfo.ImageUrl,
		&result.BusinessInfo.Info,
		&result.BusinessInfo.Email,
		&result.BusinessInfo.PhoneNumber,
		&result.BusinessInfo.WebSite,
		&result.BusinessInfo.TelegramAccount,
		&result.BusinessInfo.InstagramAccount,
		&result.BusinessInfo.LinkedInAccount,
		&result.BusinessInfo.CreatedAt,
		&result.ViewsCount,
		&result.CreatedAt,
	); err != nil {
		return &repo.VacancyResponse{}, err
	}
	return &result, nil

}

func (vr *vacancyRepo) CreateVacancy(req *repo.CreateVacancyReq) (*repo.VacancyResponse, error) {
	query := `
		INSERT INTO vacancies(
			name,
			address,
			job_type,
			info,
			min_salary,
			max_salary,
			business_id,
			category_id
		)VALUES($1,$2,$3,$4,$5,$6,$7,2)
		RETURNING
			id
	`
	var VacancyId int64
	row := vr.db.QueryRow(
		query,
		req.Name,
		utils.NullString(req.Address),
		req.JobType,
		utils.NullString(req.Info),
		utils.NullFloat64(req.MinSalary),
		utils.NullFloat64(req.MaxSalary),
		req.BusinessId,
	)
	if err := row.Scan(&VacancyId); err != nil {
		return &repo.VacancyResponse{}, err
	}
	return vr.GetVacancy(VacancyId)
}

func (vr *vacancyRepo) GetVacancy(req int64) (*repo.VacancyResponse, error) {

	query := `
		SELECT
			v.id,
			v.name,
			c.id, 
			c.name,
			COALESCE(v.image_url,'')as image_url,
			COALESCE(v.address,'') as address,
			v.job_type,
			COALESCE(v.info,'') as info,
			COALESCE(v.min_salary,0) as min_salary,
			COALESCE(v.max_salary,0) as max_salary,
			b.id,
			b.name,
			COALESCE(b.address,'') as address,
			COALESCE(b.image_url,'') as image_url,
			COALESCE(b.info,'') as info,
			b.email,
			COALESCE(b.phone_number,'') as phone_number,
			COALESCE(b.web_site,'') as web_site,
			COALESCE(b.telegram_account,'') as telegram_account,
			COALESCE(b.instagram_account,'') as instagram_account,
			COALESCE(b.linked_in_account,'') as linked_in_account,
			b.created_at,
			COALESCE(v.views_count,1) as views_count,
			v.created_at
		FROM vacancies v
		INNER JOIN businesses b
			ON b.id=v.business_id
		INNER JOIN categories c
			on c.id=v.category_id
		WHERE v.id=$1
	`
	row := vr.db.QueryRow(query, req)
	return ConvertVacancy(row)
}

func (vr *vacancyRepo) GetAllVacancies(params *repo.GetAllVacanciesReq) (*repo.GetAllVacanciesRes, error) {
	result := repo.GetAllVacanciesRes{
		Vacancies: make([]*repo.VacancyResponse, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " WHERE true "
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			and (v.name ILIKE '%s' OR c.name ILIKE '%s' OR v.image_url ILIKE '%s'
		OR v.address ILIKE '%s' OR v.job_type ILIKE '%s' OR v.info ILIKE '%s' 
		OR b.web_site ILIKE '%s' OR b.telegram_account ILIKE '%s'
		OR b.instagram_account ILIKE '%s' OR b.linked_in_account ILIKE '%s'
		OR b.name ILIKE '%s' OR b.address ILIKE '%s'
		OR b.image_url ILIKE '%s' OR b.info ILIKE '%s'
		OR b.email ILIKE '%s' OR b.phone_number ILIKE '%s')`, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str)
	}

	if params.SortByDate == "" {
		params.SortByDate = "desc"
	}

	query := `
		SELECT
			v.id,
			v.name,
			c.id, 
			c.name,
			COALESCE(v.image_url,'')as image_url,
			COALESCE(v.address,'') as address,
			v.job_type,
			COALESCE(v.info,'') as info,
			COALESCE(v.min_salary,0) as min_salary,
			COALESCE(v.max_salary,0) as max_salary,
			b.id,
			b.name,
			COALESCE(b.address,'') as address,
			COALESCE(b.image_url,'') as image_url,
			COALESCE(b.info,'') as info,
			b.email,
			COALESCE(b.phone_number,'') as phone_number,
			COALESCE(b.web_site,'') as web_site,
			COALESCE(b.telegram_account,'') as telegram_account,
			COALESCE(b.instagram_account,'') as instagram_account,
			COALESCE(b.linked_in_account,'') as linked_in_account,
			b.created_at,
			COALESCE(v.views_count,1) as views_count,
			v.created_at
		FROM vacancies v
		INNER JOIN businesses b
			ON b.id=v.business_id
		INNER JOIN categories c
			ON c.id=v.category_id
	` + filter + ` ORDER BY v.created_at ` + params.SortByDate + ` ` + limit
	rows, err := vr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var vacancy repo.VacancyResponse
		if err := rows.Scan(
			&vacancy.Id,
			&vacancy.Name,
			&vacancy.CategoryInfo.Id,
			&vacancy.CategoryInfo.Name,
			&vacancy.ImageUrl,
			&vacancy.Address,
			&vacancy.JobType,
			&vacancy.Info,
			&vacancy.MinSalary,
			&vacancy.MaxSalary,
			&vacancy.BusinessInfo.Id,
			&vacancy.BusinessInfo.Name,
			&vacancy.BusinessInfo.Address,
			&vacancy.BusinessInfo.ImageUrl,
			&vacancy.BusinessInfo.Info,
			&vacancy.BusinessInfo.Email,
			&vacancy.BusinessInfo.PhoneNumber,
			&vacancy.BusinessInfo.WebSite,
			&vacancy.BusinessInfo.TelegramAccount,
			&vacancy.BusinessInfo.InstagramAccount,
			&vacancy.BusinessInfo.LinkedInAccount,
			&vacancy.BusinessInfo.CreatedAt,
			&vacancy.ViewsCount,
			&vacancy.CreatedAt,
		); err != nil {
			return &repo.GetAllVacanciesRes{}, err
		}
		result.Vacancies = append(result.Vacancies, &vacancy)
	}
	queryCount := `SELECT count(1) FROM vacancies ` + filter
	err = vr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (vr *vacancyRepo) UpdateVacancy(req *repo.UpdateVacancyReq) (*repo.VacancyResponse, error) {
	query := `
		UPDATE vacancies SET
			name=$1,
			job_type=$2,
			info=$3,
			min_salary=$4,
			max_salary=$5
		where id=$6
	`
	effect, err := vr.db.Exec(
		query,
		req.Name,
		req.JobType,
		req.Info,
		req.MinSalary,
		req.MaxSalary,
		req.Id,
	)
	if err != nil {
		return &repo.VacancyResponse{}, err
	}
	rows, err := effect.RowsAffected()
	if err != nil {
		return &repo.VacancyResponse{}, err
	}
	if rows == 0 {
		return &repo.VacancyResponse{}, sql.ErrNoRows
	}
	return vr.GetVacancy(req.Id)
}

func (vr *vacancyRepo) DeleteVacancy(req int64) error {
	effect, err := vr.db.Exec("delete from vacancies where id=$1", req)
	if err != nil {
		return err
	}
	rows, err := effect.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (vr *vacancyRepo) UploadVacancyImage(req *repo.UploadVacancyImageReq) (*repo.VacancyResponse, error) {
	query := `UPDATE vacancies SET
		image_url=$1
	where id = $2
	`
	effect, err := vr.db.Exec(query, req.ImageUrl, req.VacancyId)
	if err != nil {
		return &repo.VacancyResponse{}, err
	}
	rows, err := effect.RowsAffected()
	if err != nil {
		return &repo.VacancyResponse{}, err
	}
	if rows == 0 {
		return &repo.VacancyResponse{}, sql.ErrNoRows
	}
	return vr.GetVacancy(req.VacancyId)
}
