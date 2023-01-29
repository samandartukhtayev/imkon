package repo

import "time"

type User struct {
	Id           int64
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Password     string
	ImageUrl     string
	PortfoliaUrl string
	CreatedAt    time.Time
}

type CreateUserReq struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
}

type GetAllUsersReq struct {
	Page       int64
	Limit      int64
	Search     string
	SortByDate string
}

// GetAllAcceptedUsers

type GetAllAcceptedUsersReq struct {
	Page       int64
	Limit      int64
	Search     string
	VacancyId  int64
	SortByDate string
}

type GetAllUsersResp struct {
	Users []*User
	Count int64
}

type UpdateUserReq struct {
	Id          int64
	FirstName   string
	LastName    string
	PhoneNumber string
}

type UploadUserImageReq struct {
	UserId   int64
	ImageUrl string
}

type UploadUserPortfoliaReq struct {
	UserId       int64
	PortfoliaUrl string
}

type UserStorageI interface {
	CreateUser(*CreateUserReq) (*User, error)
	GetUser(int64) (*User, error)
	GetAllUsers(*GetAllUsersReq) (*GetAllUsersResp, error)
	GetAllAcceptedUsers(*GetAllAcceptedUsersReq) (*GetAllUsersResp, error)
	UpdateUser(*UpdateUserReq) (*User, error)
	DeleteUser(int64) error
	UploadUserImage(*UploadUserImageReq) (*User, error)
	UploadUserPortfolia(*UploadUserPortfoliaReq) (*User, error)
}
