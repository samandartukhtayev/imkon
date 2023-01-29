package repo

import "time"

type Business struct {
	Id               int64
	Name             string
	Password         string
	Address          string
	ImageUrl         string
	Info             string
	Email            string
	PhoneNumber      string
	WebSite          string
	TelegramAccount  string
	InstagramAccount string
	LinkedInAccount  string
	CreatedAt        time.Time
}

type CreateBusinessReq struct {
	Name             string
	Password         string
	Address          string
	Info             string
	Email            string
	PhoneNumber      string
	WebSite          string
	TelegramAccount  string
	InstagramAccount string
	LinkedInAccount  string
}

type UpdateBusinessReq struct {
	Id               int64
	Name             string
	Address          string
	Info             string
	PhoneNumber      string
	WebSite          string
	TelegramAccount  string
	InstagramAccount string
	LinkedInAccount  string
}

type BusinessResponse struct {
	Id               int64
	Name             string
	Address          string
	ImageUrl         string
	Info             string
	Email            string
	PhoneNumber      string
	WebSite          string
	TelegramAccount  string
	InstagramAccount string
	LinkedInAccount  string
	CreatedAt        time.Time
}

type UploadBusinessImageReq struct {
	BusinessId int64
	ImageUrl   string
}

type GetAllBusinessesReq struct {
	Page       int64
	Limit      int64
	Search     string
	SortByDate string
}

type GetAllBusinessesRes struct {
	Businesses []*BusinessResponse
	Count      int64
}

type BusinessStorageI interface {
	CreateBusiness(*CreateBusinessReq) (*BusinessResponse, error)
	GetBusiness(int64) (*BusinessResponse, error)
	GetAllBusinesses(*GetAllBusinessesReq) (*GetAllBusinessesRes, error)
	UpdateBusiness(*UpdateBusinessReq) (*BusinessResponse, error)
	DeleteBusiness(int64) error
	UploadBusinessImage(*UploadBusinessImageReq) (*BusinessResponse, error)
}
