package models

import "time"

type Business struct {
	Id               int64     `json:"id"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	Address          string    `json:"address"`
	ImageUrl         string    `json:"image_url"`
	Info             string    `json:"info"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	WebSite          string    `json:"web_site"`
	TelegramAccount  string    `json:"telegram_account"`
	InstagramAccount string    `json:"instagram_account"`
	LinkedInAccount  string    `json:"linked_in_account"`
	CreatedAt        time.Time `json:"created_at"`
}

type CreateBusinessReq struct {
	Name             string `json:"name"`
	Password         string `json:"password"`
	Address          string `json:"address"`
	Info             string `json:"info"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	WebSite          string `json:"web_site"`
	TelegramAccount  string `json:"telegram_account"`
	InstagramAccount string `json:"instagram_account"`
	LinkedInAccount  string `json:"linked_in_account"`
}

type UpdateBusinessReq struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Info             string `json:"info"`
	PhoneNumber      string `json:"phone_number"`
	WebSite          string `json:"web_site"`
	TelegramAccount  string `json:"telegram_account"`
	InstagramAccount string `json:"instagram_account"`
	LinkedInAccount  string `json:"linked_in_account"`
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
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllBusinessesResp struct {
	Businesses []*BusinessResponse `json:"businesses"`
	Count      int64               `json:"count"`
}
