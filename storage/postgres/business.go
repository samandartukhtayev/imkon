package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/pkg/utils"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type businessRepo struct {
	db *sqlx.DB
}

func NewBusiness(db *sqlx.DB) repo.BusinessStorageI {
	return &businessRepo{db: db}
}

func ConvertBusiness(row *sql.Row) (*repo.BusinessResponse, error) {
	var result repo.BusinessResponse
	if err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Address,
		&result.ImageUrl,
		&result.Info,
		&result.Email,
		&result.PhoneNumber,
		&result.WebSite,
		&result.TelegramAccount,
		&result.InstagramAccount,
		&result.LinkedInAccount,
		&result.CreatedAt,
	); err != nil {
		return &repo.BusinessResponse{}, err
	}
	return &result, nil
}

func (br *businessRepo) CreateBusiness(req *repo.CreateBusinessReq) (*repo.BusinessResponse, error) {
	query := `
		INSERT INTO businesses(
			name,
			password,
			address,
			info,
			email,
			phone_number,
			web_site,
			telegram_account,
			instagram_account,
			linked_in_account
		)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING
			id,
			name,
			COALESCE(address,'') as address,
			COALESCE(image_url,'') as image_url,
			COALESCE(info,'') as info,
			email,
			COALESCE(phone_number,'') as phone_number,
			COALESCE(web_site,'') as web_site,
			COALESCE(telegram_account,'') as telegram_account,
			COALESCE(instagram_account,'') as instagram_account,
			COALESCE(linked_in_account,'') as linked_in_account,
			created_at
	`
	row := br.db.QueryRow(
		query,
		req.Name,
		req.Password,
		utils.NullString(req.Address),
		utils.NullString(req.Info),
		req.Email,
		utils.NullString(req.PhoneNumber),
		utils.NullString(req.WebSite),
		utils.NullString(req.TelegramAccount),
		utils.NullString(req.InstagramAccount),
		utils.NullString(req.LinkedInAccount),
	)
	return ConvertBusiness(row)
}

func (br *businessRepo) GetBusiness(req int64) (*repo.BusinessResponse, error) {
	query := `
		SELECT 
			id,
			name,
			COALESCE(address,'') as address,
			COALESCE(image_url,'') as image_url,
			COALESCE(info,'') as info,
			email,
			COALESCE(phone_number,'') as phone_number,
			COALESCE(web_site,'') as web_site,
			COALESCE(telegram_account,'') as telegram_account,
			COALESCE(instagram_account,'') as instagram_account,
			COALESCE(linked_in_account,'') as linked_in_account,
			created_at
		from businesses
		where id=$1
	`
	row := br.db.QueryRow(query, req)
	return ConvertBusiness(row)
}

func (br *businessRepo) GetAllBusinesses(params *repo.GetAllBusinessesReq) (*repo.GetAllBusinessesRes, error) {
	result := repo.GetAllBusinessesRes{
		Businesses: make([]*repo.BusinessResponse, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)
	filter := " WHERE true "
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			and (name ILIKE '%s' OR address ILIKE '%s' OR image_url ILIKE '%s'
		OR info ILIKE '%s' OR email ILIKE '%s' OR phone_number ILIKE '%s' 
		OR web_site ILIKE '%s' OR telegram_account ILIKE '%s'
		OR instagram_account ILIKE '%s' OR linked_in_account ILIKE '%s')`, str, str, str, str, str, str, str, str, str, str)
	}

	if params.SortByDate == "" {
		params.SortByDate = "desc"
	}
	query := `
		SELECT
			id,
			name,
			COALESCE(address,'') as address,
			COALESCE(image_url,'') as image_url,
			COALESCE(info,'') as info,
			email,
			COALESCE(phone_number,'') as phone_number,
			COALESCE(web_site,'') as web_site,
			COALESCE(telegram_account,'') as telegram_account,
			COALESCE(instagram_account,'') as instagram_account,
			COALESCE(linked_in_account,'') as linked_in_account,
			created_at
		FROM businesses ` + filter + ` ORDER BY created_at ` + params.SortByDate + ` ` + limit

	rows, err := br.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var business repo.BusinessResponse
		if err := rows.Scan(
			&business.Id,
			&business.Name,
			&business.Address,
			&business.ImageUrl,
			&business.Info,
			&business.Email,
			&business.PhoneNumber,
			&business.WebSite,
			&business.TelegramAccount,
			&business.InstagramAccount,
			&business.LinkedInAccount,
			&business.CreatedAt,
		); err != nil {
			return &repo.GetAllBusinessesRes{}, err
		}
		result.Businesses = append(result.Businesses, &business)
	}
	queryCount := `SELECT count(1) FROM businesses ` + filter
	err = br.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (br *businessRepo) UpdateBusiness(req *repo.UpdateBusinessReq) (*repo.BusinessResponse, error) {
	query := `
		UPDATE businesses SET
			name=$1,
			address=$2,
			info=$3,
			phone_number=$4,
			web_site=$5,
			telegram_account=$6,
			instagram_account=$7,
			linked_in_account=$8
		where id=$9
		RETURNING
			id,
			name,
			COALESCE(address,'') as address,
			COALESCE(image_url,'') as image_url,
			COALESCE(info,'') as info,
			email,
			COALESCE(phone_number,'') as phone_number,
			COALESCE(web_site,'') as web_site,
			COALESCE(telegram_account,'') as telegram_account,
			COALESCE(instagram_account,'') as instagram_account,
			COALESCE(linked_in_account,'') as linked_in_account,
			created_at
	`
	row := br.db.QueryRow(
		query,
		req.Name,
		req.Address,
		req.Info,
		req.PhoneNumber,
		req.WebSite,
		req.TelegramAccount,
		req.InstagramAccount,
		req.LinkedInAccount,
		req.Id,
	)
	return ConvertBusiness(row)
}

func (br *businessRepo) DeleteBusiness(req int64) error {
	effect, err := br.db.Exec("delete from businesses where id=$1", req)
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

func (ur *businessRepo) UploadBusinessImage(req *repo.UploadBusinessImageReq) (*repo.BusinessResponse, error) {
	query := `UPDATE businesses SET
		image_url=$1
	where id = $2
	RETURNING
		id,
		name,
		COALESCE(address,'') as address,
		COALESCE(image_url,'') as image_url,
		COALESCE(info,'') as info,
		email,
		COALESCE(phone_number,'') as phone_number,
		COALESCE(web_site,'') as web_site,
		COALESCE(telegram_account,'') as telegram_account,
		COALESCE(instagram_account,'') as instagram_account,
		COALESCE(linked_in_account,'') as linked_in_account,
		created_at
	`
	fmt.Println("----------------------------")
	fmt.Println(req.ImageUrl, req.BusinessId)
	fmt.Println("----------------------------")
	row := ur.db.QueryRow(query, req.ImageUrl, req.BusinessId)
	return ConvertBusiness(row)
}
