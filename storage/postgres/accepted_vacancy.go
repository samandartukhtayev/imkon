package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type acceptedVacancyRepo struct {
	db *sqlx.DB
}

func NewAcceptedVacancy(db *sqlx.DB) repo.AcceptedVacancyStorageI {
	return &acceptedVacancyRepo{db: db}
}

func (avr *acceptedVacancyRepo) AcceptVacancy(req *repo.AcceptVacancyReq) (*repo.AcceptVacancyRes, error) {
	var AcceptVacancyId int64
	query := `
		INSERT INTO accepted_vacancies(
			user_id,
			vacancy_id
		)VALUES($1,$2)
		RETURING
			id
	`
	row := avr.db.QueryRow(
		query,
		req.UserId,
		req.VacancyId,
	)
	if err := row.Scan(
		&AcceptVacancyId,
	); err != nil {
		return &repo.AcceptVacancyRes{}, err
	}

	return avr.GetAcceptedVacanciesById(AcceptVacancyId)
}

func (avr *acceptedVacancyRepo) GetAcceptedVacanciesByUserId(UserId int64) (*repo.AcceptedVacanciesResponse, error) {
	var result repo.AcceptedVacanciesResponse
	query := `
		SELECT 
			av.id,
			u.id,
			u.first_name,
			COALESCE(u.last_name,'') as last_name,
			u.email,
			COALESCE(u.phone_number,'')as phone_number,
			COALESCE(u.image_url,'') as image_url,
			COALESCE(u.portfolia_url,'') as portfolia_url,
			u.created_at
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
			b.created_at
			av.created_at
		FROM accepted_vacancies av 
		INNER JOIN users u
			ON av.user_id=u.id
		INNER JOIN businesses b
			ON av.business_id=b.id
		WHERE av.user_id=$1
	`
	rows, err := avr.db.Query(
		query,
		UserId,
	)
	if err != nil {
		return &repo.AcceptedVacanciesResponse{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var acceptedVacancy repo.AcceptVacancyRes
		if err = rows.Scan(
			&acceptedVacancy.Id,
			&acceptedVacancy.UserInfo.Id,
			&acceptedVacancy.UserInfo.FirstName,
			&acceptedVacancy.UserInfo.LastName,
			&acceptedVacancy.UserInfo.Email,
			&acceptedVacancy.UserInfo.PhoneNumber,
			&acceptedVacancy.UserInfo.ImageUrl,
			&acceptedVacancy.UserInfo.PortfoliaUrl,
			&acceptedVacancy.UserInfo.CreatedAt,
			&acceptedVacancy.BusinessInfo.Id,
			&acceptedVacancy.BusinessInfo.Name,
			&acceptedVacancy.BusinessInfo.Address,
			&acceptedVacancy.BusinessInfo.ImageUrl,
			&acceptedVacancy.BusinessInfo.Info,
			&acceptedVacancy.BusinessInfo.Email,
			&acceptedVacancy.BusinessInfo.PhoneNumber,
			&acceptedVacancy.BusinessInfo.WebSite,
			&acceptedVacancy.BusinessInfo.WebSite,
			&acceptedVacancy.BusinessInfo.TelegramAccount,
			&acceptedVacancy.BusinessInfo.InstagramAccount,
			&acceptedVacancy.BusinessInfo.LinkedInAccount,
			&acceptedVacancy.BusinessInfo.CreatedAt,
			&acceptedVacancy.CreatedAt,
		); err != nil {
			return &repo.AcceptedVacanciesResponse{}, err
		}
		result.AcceptedVacancies = append(result.AcceptedVacancies, &acceptedVacancy)
		result.Count++
	}
	return &result, nil
}
func (avr *acceptedVacancyRepo) GetAcceptedVacanciesByBusinessId(BusinessId int64) (*repo.AcceptedVacanciesResponse, error) {
	var result repo.AcceptedVacanciesResponse
	query := `
		SELECT 
			av.id,
			u.id,
			u.first_name,
			COALESCE(u.last_name,'') as last_name,
			u.email,
			COALESCE(u.phone_number,'')as phone_number,
			COALESCE(u.image_url,'') as image_url,
			COALESCE(u.portfolia_url,'') as portfolia_url,
			u.created_at
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
			b.created_at
			av.created_at
		FROM accepted_vacancies av 
		INNER JOIN users u
			ON av.user_id=u.id
		INNER JOIN businesses b
			ON av.business_id=b.id
		WHERE av.business_id=$1
	`
	rows, err := avr.db.Query(
		query,
		BusinessId,
	)
	if err != nil {
		return &repo.AcceptedVacanciesResponse{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var acceptedVacancy repo.AcceptVacancyRes
		if err = rows.Scan(
			&acceptedVacancy.Id,
			&acceptedVacancy.UserInfo.Id,
			&acceptedVacancy.UserInfo.FirstName,
			&acceptedVacancy.UserInfo.LastName,
			&acceptedVacancy.UserInfo.Email,
			&acceptedVacancy.UserInfo.PhoneNumber,
			&acceptedVacancy.UserInfo.ImageUrl,
			&acceptedVacancy.UserInfo.PortfoliaUrl,
			&acceptedVacancy.UserInfo.CreatedAt,
			&acceptedVacancy.BusinessInfo.Id,
			&acceptedVacancy.BusinessInfo.Name,
			&acceptedVacancy.BusinessInfo.Address,
			&acceptedVacancy.BusinessInfo.ImageUrl,
			&acceptedVacancy.BusinessInfo.Info,
			&acceptedVacancy.BusinessInfo.Email,
			&acceptedVacancy.BusinessInfo.PhoneNumber,
			&acceptedVacancy.BusinessInfo.WebSite,
			&acceptedVacancy.BusinessInfo.WebSite,
			&acceptedVacancy.BusinessInfo.TelegramAccount,
			&acceptedVacancy.BusinessInfo.InstagramAccount,
			&acceptedVacancy.BusinessInfo.LinkedInAccount,
			&acceptedVacancy.BusinessInfo.CreatedAt,
			&acceptedVacancy.CreatedAt,
		); err != nil {
			return &repo.AcceptedVacanciesResponse{}, err
		}
		result.AcceptedVacancies = append(result.AcceptedVacancies, &acceptedVacancy)
		result.Count++
	}
	return &result, nil
}

func (avr *acceptedVacancyRepo) GetAcceptedVacanciesById(Id int64) (*repo.AcceptVacancyRes, error) {
	var result repo.AcceptVacancyRes
	query := `
		SELECT 
			av.id,
			u.id,
			u.first_name,
			COALESCE(u.last_name,'') as last_name,
			u.email,
			COALESCE(u.phone_number,'')as phone_number,
			COALESCE(u.image_url,'') as image_url,
			COALESCE(u.portfolia_url,'') as portfolia_url,
			u.created_at
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
			b.created_at
			av.created_at
		FROM accepted_vacancies av 
		INNER JOIN users u
			ON av.user_id=u.id
		INNER JOIN businesses b
			ON av.business_id=b.id
		WHERE av.id=$1
	`
	row := avr.db.QueryRow(
		query,
		Id,
	)
	if err := row.Scan(
		&result.Id,
		&result.UserInfo.Id,
		&result.UserInfo.FirstName,
		&result.UserInfo.LastName,
		&result.UserInfo.Email,
		&result.UserInfo.PhoneNumber,
		&result.UserInfo.ImageUrl,
		&result.UserInfo.PortfoliaUrl,
		&result.UserInfo.CreatedAt,
		&result.BusinessInfo.Id,
		&result.BusinessInfo.Name,
		&result.BusinessInfo.Address,
		&result.BusinessInfo.ImageUrl,
		&result.BusinessInfo.Info,
		&result.BusinessInfo.Email,
		&result.BusinessInfo.PhoneNumber,
		&result.BusinessInfo.WebSite,
		&result.BusinessInfo.WebSite,
		&result.BusinessInfo.TelegramAccount,
		&result.BusinessInfo.InstagramAccount,
		&result.BusinessInfo.LinkedInAccount,
		&result.BusinessInfo.CreatedAt,
		&result.CreatedAt,
	); err != nil {
		return &repo.AcceptVacancyRes{}, err
	}
	return &result, nil
}

func (avr *acceptedVacancyRepo) DeleteAcceptedVacancy(Id int64) error {
	effect, err := avr.db.Exec(`DELETE FROM accepted_vacancies WHERE id=$1`, Id)
	if err != nil {
		return err
	}
	if count, _ := effect.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
