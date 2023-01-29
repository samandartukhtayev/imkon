package repo

import "time"

type AcceptVacancyStruct struct {
	Id        int64
	UserId    int64
	VacancyId int64
	CreatedAt time.Time
}

type AcceptVacancyReq struct {
	UserId    int64
	VacancyId int64
}

type AcceptVacancyRes struct {
	Id           int64
	UserInfo     User
	BusinessInfo BusinessResponse
	CreatedAt    time.Time
}

type AcceptedVacanciesResponse struct {
	AcceptedVacancies []*AcceptVacancyRes
	Count             int64
}

type AcceptedVacancyStorageI interface {
	AcceptVacancy(*AcceptVacancyReq) (*AcceptVacancyRes, error)
	GetAcceptedVacanciesByUserId(int64) (*AcceptedVacanciesResponse, error)
	GetAcceptedVacanciesById(int64) (*AcceptVacancyRes, error)
	GetAcceptedVacanciesByBusinessId(int64) (*AcceptedVacanciesResponse, error)
	DeleteAcceptedVacancy(Id int64) error
}
