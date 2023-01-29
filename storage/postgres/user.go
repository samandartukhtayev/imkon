package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/pkg/utils"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}

func ConvertUser(row *sql.Row) (*repo.User, error) {
	var result repo.User
	if err := row.Scan(
		&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PhoneNumber,
		&result.ImageUrl,
		&result.PortfoliaUrl,
		&result.CreatedAt,
	); err != nil {
		return &repo.User{}, err
	}
	return &result, nil
}

func (ur *userRepo) CreateUser(req *repo.CreateUserReq) (*repo.User, error) {
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			email,
			phone_number,
			password
		)VALUES($1,$2,$3,$4,$5)
		RETURNING
			id,
			first_name,
			COALESCE(last_name,'') as last_name,
			email,
			COALESCE(phone_number,'')as phone_number,
			COALESCE(image_url,'') as image_url,
			COALESCE(portfolia_url,'') as portfolia_url,
			created_at
	`
	row := ur.db.QueryRow(
		query,
		req.FirstName,
		utils.NullString(req.LastName),
		req.Email,
		utils.NullString(req.PhoneNumber),
		utils.NullString(req.Password),
	)
	return ConvertUser(row)
}

func (ur *userRepo) GetUser(req int64) (*repo.User, error) {
	query := `
		SELECT 
			id,
			first_name,
			COALESCE(last_name,'') as last_name,
			email,
			COALESCE(phone_number,'')as phone_number,
			COALESCE(image_url,'') as image_url,
			COALESCE(portfolia_url,'') as portfolia_url,
			created_at
		from users
		WHERE id=$1
	`
	row := ur.db.QueryRow(query, req)
	return ConvertUser(row)
}

func (ur *userRepo) GetAllUsers(req *repo.GetAllUsersReq) (*repo.GetAllUsersResp, error) {
	result := repo.GetAllUsersResp{
		Users: make([]*repo.User, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := " WHERE true "
	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			and (first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s'
		OR image_url ILIKE '%s' OR phone_number ILIKE '%s' OR portfolia_url ILIKE '%s')`, str, str, str, str, str, str)
	}

	if req.SortByDate == "" {
		req.SortByDate = "desc"
	}
	query := `
		SELECT
			id,
			first_name,
			COALESCE(last_name,'') as last_name,
			email,
			COALESCE(phone_number,'')as phone_number,
			COALESCE(image_url,'') as image_url,
			COALESCE(portfolia_url,'') as portfolia_url,
			created_at
		FROM users ` + filter + ` ORDER BY created_at ` + req.SortByDate + ` ` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var usr repo.User
		if err := rows.Scan(
			&usr.Id,
			&usr.FirstName,
			&usr.LastName,
			&usr.Email,
			&usr.PhoneNumber,
			&usr.ImageUrl,
			&usr.PortfoliaUrl,
			&usr.CreatedAt,
		); err != nil {
			return &repo.GetAllUsersResp{}, err
		}
		result.Users = append(result.Users, &usr)
	}
	queryCount := `SELECT count(1) FROM users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) GetAllAcceptedUsers(req *repo.GetAllAcceptedUsersReq) (*repo.GetAllUsersResp, error) {
	result := repo.GetAllUsersResp{
		Users: make([]*repo.User, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := fmt.Sprintf(` 
		WHERE id=(SELECT 
			user_id 
		FROM accepted_vacancies 
		WHERE business_id=(SELECT 
			business_id 
		FROM vacancies 
		WHERE id='%d')) `, req.VacancyId)

	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			and (first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s'
		OR image_url ILIKE '%s' OR phone_number ILIKE '%s' OR portfolia_url ILIKE '%s')`, str, str, str, str, str, str)
	}

	if req.SortByDate == "" {
		req.SortByDate = "desc"
	}
	query := `
		SELECT
			id,
			first_name,
			COALESCE(last_name,'') as last_name,
			email,
			COALESCE(phone_number,'')as phone_number,
			COALESCE(image_url,'') as image_url,
			COALESCE(portfolia_url,'') as portfolia_url,
			created_at
		FROM users ` + filter + ` ORDER BY created_at ` + req.SortByDate + ` ` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var usr repo.User
		if err := rows.Scan(
			&usr.Id,
			&usr.FirstName,
			&usr.LastName,
			&usr.Email,
			&usr.PhoneNumber,
			&usr.ImageUrl,
			&usr.PortfoliaUrl,
			&usr.CreatedAt,
		); err != nil {
			return &repo.GetAllUsersResp{}, err
		}
		result.Users = append(result.Users, &usr)
	}

	queryCount := `SELECT count(1) FROM users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) UpdateUser(req *repo.UpdateUserReq) (*repo.User, error) {
	query := `
		UPDATE users SET
			first_name = $1,
			last_name = $2,
			phone_number = $3
		where id = $4
		RETURNING
			id,
			first_name,
			COALESCE(last_name,'') as last_name,
			email,
			COALESCE(phone_number,'')as phone_number,
			COALESCE(image_url,'') as image_url,
			COALESCE(portfolia_url,'') as portfolia_url,
			created_at
	`
	row := ur.db.QueryRow(query, req.FirstName, req.LastName, req.PhoneNumber, req.Id)
	return ConvertUser(row)
}

func (ur *userRepo) DeleteUser(req int64) error {
	effect, err := ur.db.Exec("delete from users where id=$1", req)
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

func (ur *userRepo) UploadUserImage(req *repo.UploadUserImageReq) (*repo.User, error) {
	query := `UPDATE users SET
		image_url=$1
	where id = $2
	RETURNING
		id,
		first_name,
		COALESCE(last_name,'') as last_name,
		email,
		COALESCE(phone_number,'')as phone_number,
		COALESCE(image_url,'') as image_url,
		COALESCE(portfolia_url,'') as portfolia_url,
		created_at
	`
	row := ur.db.QueryRow(query, req.ImageUrl, req.UserId)
	return ConvertUser(row)
}

func (ur *userRepo) UploadUserPortfolia(req *repo.UploadUserPortfoliaReq) (*repo.User, error) {
	query := `UPDATE users SET
		portfolia_url=$1
	where id = $2
	RETURNING
		id,
		first_name,
		COALESCE(last_name,'') as last_name,
		email,
		COALESCE(phone_number,'')as phone_number,
		COALESCE(image_url,'') as image_url,
		COALESCE(portfolia_url,'') as portfolia_url,
		created_at
	`
	row := ur.db.QueryRow(query, req.PortfoliaUrl, req.UserId)
	return ConvertUser(row)
}
