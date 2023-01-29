package models

import "time"

type Vacancy struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	CategoryId int64     `json:"category_id"`
	ImageUrl   string    `json:"image_url"`
	Address    string    `json:"address"`
	JobType    string    `json:"job_type"`
	MinSalary  float64   `json:"min_salary"`
	MaxSalary  float64   `json:"max_salary"`
	Info       string    `json:"info"`
	ViewsCount int64     `json:"views_count"`
	BusinessId int64     `json:"business_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateVacancyReq struct {
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	JobType    string  `json:"job_type"`
	MinSalary  float64 `json:"min_salary"`
	MaxSalary  float64 `json:"max_salary"`
	Info       string  `json:"info"`
	BusinessId int64   `json:"business_id"`
}

type VacancyResponse struct {
	Id           int64            `json:"id"`
	Name         string           `json:"name"`
	CategoryInfo Category         `json:"category_info"`
	ImageUrl     string           `json:"image_url"`
	Address      string           `json:"address"`
	JobType      string           `json:"job_type"`
	Info         string           `json:"info"`
	MinSalary    float64          `json:"min_salary"`
	MaxSalary    float64          `json:"max_salary"`
	BusinessInfo BusinessResponse `json:"business_info"`
	ViewsCount   int64            `json:"views_count"`
	CreatedAt    time.Time        `json:"created_at"`
}

type UpdateVacancyReq struct {
	Name      string  `json:"name"`
	JobType   string  `json:"job_type"`
	Info      string  `json:"info"`
	MinSalary float64 `json:"min_salary"`
	MaxSalary float64 `json:"max_salary"`
}
type UploadVacancyImageReq struct {
	VacancyId int64  `json:"vacancy_id"`
	ImageUrl  string `json:"image_url"`
}

type GetAllVacanciesReq struct {
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}
type GetAllVacanciesRes struct {
	Vacancies []*VacancyResponse `json:"vacancies"`
	Count     int64              `json:"count"`
}
