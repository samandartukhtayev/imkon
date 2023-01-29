package models

import "time"

type AcceptVacancyStruct struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	VacancyId int64     `json:"vacancy_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AcceptVacancyReq struct {
	UserId    int64 `json:"user_id"`
	VacancyId int64 `json:"vacancy_id"`
}

type AcceptVacancyRes struct {
	Id           int64            `json:"id"`
	UserInfo     UserResponse     `json:"user_info"`
	BusinessInfo BusinessResponse `json:"business_info"`
	CreatedAt    time.Time        `json:"created_at"`
}

type AcceptedVacanciesResponse struct {
	AcceptedVacancies []*AcceptVacancyRes `json:"accepted_vacancies"`
	Count             int64               `json:"count"`
}
