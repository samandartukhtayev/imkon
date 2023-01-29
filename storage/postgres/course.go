package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type courseRepo struct {
	db *sqlx.DB
}

func NewCourse(db *sqlx.DB) repo.CourseStorageI {
	return &courseRepo{db: db}
}

func ConvertCourse(row *sql.Row) (*repo.CourseResponse, error) {
	var result repo.CourseResponse
	if err := row.Scan(
		&result.Id,
		&result.Name,
		&result.CoursePrice,
		&result.Category.Id,
		&result.Category.Name,
		&result.Info,
		&result.Business.Id,
		&result.Business.Name,
		&result.Business.Address,
		&result.Business.ImageUrl,
		&result.Business.Info,
		&result.Business.Email,
		&result.Business.PhoneNumber,
		&result.Business.WebSite,
		&result.Business.TelegramAccount,
		&result.Business.InstagramAccount,
		&result.Business.LinkedInAccount,
		&result.Business.CreatedAt,
		&result.ImageUrl,
		&result.SaleOf,
		&result.CreatedAt,
	); err != nil {
		return &repo.CourseResponse{}, err
	}
	return &result, nil
}

func (cr *courseRepo) CreateCourse(req *repo.CreateCourseReq) (*repo.CourseResponse, error) {
	query := `
		INSERT INTO courses(
			name,
			course_price,
			info,
			business_id,
			sale_of,
			category_id
		)VALUES($1,$2,$3,$4,$5,1)
		RETURNING
		id
	`
	var CourseId int
	row := cr.db.QueryRow(
		query,
		req.Name,
		req.CoursePrice,
		req.Info,
		req.BusinessId,
		req.SaleOf,
	)

	if err := row.Scan(&CourseId); err != nil {
		return &repo.CourseResponse{}, err
	}
	return cr.GetCourse(int64(CourseId))
}

func (cr *courseRepo) GetCourse(req int64) (*repo.CourseResponse, error) {
	query := `
		SELECT
			c.id,
			c.name,
			c.course_price,
			ct.id, 
			ct.name,
			COALESCE(c.info,'') as info,
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
			COALESCE(c.image_url, '') as image_url,
			c.sale_of,
			c.created_at
		FROM courses c
		INNER JOIN categories ct
			ON c.category_id=ct.id
		INNER JOIN businesses b
			ON c.business_id=b.id
		WHERE c.id=$1
	`

	row := cr.db.QueryRow(query, req)
	return ConvertCourse(row)
}

func (cr *courseRepo) GetAllCourses(params *repo.GetAllCoursesReq) (*repo.GetAllCoursesRes, error) {
	result := repo.GetAllCoursesRes{
		Courses: make([]*repo.CourseResponse, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)
	filter := " WHERE true "

	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			and (c.name ILIKE '%s' OR ct.name ILIKE '%s' OR c.info ILIKE '%s'
		OR b.name ILIKE '%s' OR b.address ILIKE '%s' OR b.image_url ILIKE '%s' 
		OR b.info ILIKE '%s' OR b.email ILIKE '%s'
		OR b.telegram_account ILIKE '%s' OR b.instagram_account ILIKE '%s' OR b.linked_in_account ILIKE '%s'
		OR c.image_url ILIKE '%s')`, str, str, str, str, str, str, str, str, str, str, str, str)
	}

	if params.SortByDate == "" {
		params.SortByDate = "desc"
	}

	query := `
		SELECT
			c.id,
			c.name,
			c.course_price,
			ct.id, 
			ct.name,
			COALESCE(c.info,'') as info,
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
			COALESCE(c.image_url, '') as image_url,
			c.sale_of,
			c.created_at
		FROM courses c
		INNER JOIN categories ct
			ON c.category_id=ct.id
		INNER JOIN businesses b
			ON c.business_id=b.id
	` + filter + ` ORDER BY c.created_at ` + params.SortByDate + ` ` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return &repo.GetAllCoursesRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var Course repo.CourseResponse
		if err := rows.Scan(
			&Course.Id,
			&Course.Name,
			&Course.CoursePrice,
			&Course.Category.Id,
			&Course.Category.Name,
			&Course.Info,
			&Course.Business.Id,
			&Course.Business.Name,
			&Course.Business.Address,
			&Course.Business.ImageUrl,
			&Course.Business.Info,
			&Course.Business.Email,
			&Course.Business.PhoneNumber,
			&Course.Business.WebSite,
			&Course.Business.TelegramAccount,
			&Course.Business.InstagramAccount,
			&Course.Business.LinkedInAccount,
			&Course.Business.CreatedAt,
			&Course.ImageUrl,
			&Course.SaleOf,
			&Course.CreatedAt,
		); err != nil {
			return &repo.GetAllCoursesRes{}, err
		}
		result.Courses = append(result.Courses, &Course)
	}

	queryCount := `SELECT count(1) FROM courses ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *courseRepo) UpdateCourse(req *repo.UpdateCourseReq) (*repo.CourseResponse, error) {
	query := `
		UPDATE courses SET
			name=$1,
			course_price=$2,
			info=$3,
			sale_of=$4
		WHERE id=$5
	`
	effect, err := cr.db.Exec(
		query,
		req.Name,
		req.CoursePrice,
		req.Info,
		req.SaleOf,
		req.Id,
	)
	if err != nil {
		return &repo.CourseResponse{}, err
	}
	rows, err := effect.RowsAffected()
	if err != nil {
		return &repo.CourseResponse{}, err
	}
	if rows == 0 {
		return &repo.CourseResponse{}, sql.ErrNoRows
	}

	return cr.GetCourse(req.Id)
}

func (cr *courseRepo) DeleteCourse(req int64) error {
	effect, err := cr.db.Exec("delete from courses where id=$1", req)
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

func (cr *courseRepo) UploadCourseImage(req *repo.UploadCourseImageReq) (*repo.CourseResponse, error) {
	query := `UPDATE courses SET
		image_url=$1
	where id = $2
	`
	effect, err := cr.db.Exec(query, req.ImageUrl, req.CourseId)
	if err != nil {
		return &repo.CourseResponse{}, err
	}
	rows, err := effect.RowsAffected()
	if err != nil {
		return &repo.CourseResponse{}, err
	}
	if rows == 0 {
		return &repo.CourseResponse{}, sql.ErrNoRows
	}

	return cr.GetCourse(req.CourseId)
}
