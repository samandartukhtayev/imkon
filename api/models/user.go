package models

import "time"

type User struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"password"`
	ImageUrl     string    `json:"image_url"`
	PortfoliaUrl string    `json:"portfolia_url"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserResponse struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	ImageUrl     string    `json:"image_url"`
	PortfoliaUrl string    `json:"portfolia_url"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateUserReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type GetAllUsersReq struct {
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllAcceptedUsersReq struct {
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	VacancyId  int64  `json:"vacancy_id" binding:"required"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllUsersResp struct {
	Users []*UserResponse `json:"users"`
	Count int64           `json:"count"`
}

type UpdateUserReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type UploadUserImageReq struct {
	UserId   int64  `json:"user_id"`
	ImageUrl string `json:"image_url"`
}

type UploadUserPortfoliaReq struct {
	UserId       int64  `json:"user_id"`
	PortfoliaUrl string `json:"portfolia"`
}
