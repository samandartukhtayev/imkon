package repo

import "time"

type Vacancy struct {
	Id         int64
	Name       string
	CategoryId int64
	ImageUrl   string
	Address    string
	JobType    string
	MinSalary  float64
	MaxSalary  float64
	Info       string
	ViewsCount int64
	BusinessId int64
	CreatedAt  time.Time
}

type CreateVacancyReq struct {
	Name       string
	Address    string
	JobType    string
	Info       string
	MinSalary  float64
	MaxSalary  float64
	BusinessId int64
}

type VacancyResponse struct {
	Id           int64
	Name         string
	CategoryInfo Category
	ImageUrl     string
	Address      string
	JobType      string
	Info         string
	MinSalary    float64
	MaxSalary    float64
	BusinessInfo BusinessResponse
	ViewsCount   int64
	CreatedAt    time.Time
}

type UpdateVacancyReq struct {
	Id        int64
	Name      string
	JobType   string
	Info      string
	MinSalary float64
	MaxSalary float64
}
type UploadVacancyImageReq struct {
	VacancyId int64
	ImageUrl  string
}

type GetAllVacanciesReq struct {
	Page       int64
	Limit      int64
	Search     string
	SortByDate string
}
type GetAllVacanciesRes struct {
	Vacancies []*VacancyResponse
	Count     int64
}

type VacancyStorageI interface {
	CreateVacancy(*CreateVacancyReq) (*VacancyResponse, error)
	GetVacancy(req int64) (*VacancyResponse, error)
	GetAllVacancies(*GetAllVacanciesReq) (*GetAllVacanciesRes, error)
	UpdateVacancy(*UpdateVacancyReq) (*VacancyResponse, error)
	DeleteVacancy(int64) error
	UploadVacancyImage(req *UploadVacancyImageReq) (*VacancyResponse, error)
}
