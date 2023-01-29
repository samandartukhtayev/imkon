package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandartukhtayev/imkon/storage/postgres"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

type StorageI interface {
	AcceptVacancy() repo.AcceptedVacancyStorageI
	Business() repo.BusinessStorageI
	Category() repo.CategoryStorageI
	Course() repo.CourseStorageI
	User() repo.UserStorageI
	Vacancy() repo.VacancyStorageI
}

type storagePg struct {
	userRepo          repo.UserStorageI
	acceptVacancyRepo repo.AcceptedVacancyStorageI
	businessRepo      repo.BusinessStorageI
	courseRepo        repo.CourseStorageI
	vacancyRepo       repo.VacancyStorageI
	categoryRepo      repo.CategoryStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:          postgres.NewUser(db),
		businessRepo:      postgres.NewBusiness(db),
		acceptVacancyRepo: postgres.NewAcceptedVacancy(db),
		courseRepo:        postgres.NewCourse(db),
		vacancyRepo:       postgres.NewVacancy(db),
		categoryRepo:      postgres.NewCategory(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Business() repo.BusinessStorageI {
	return s.businessRepo
}

func (s *storagePg) AcceptVacancy() repo.AcceptedVacancyStorageI {
	return s.acceptVacancyRepo
}

func (s *storagePg) Course() repo.CourseStorageI {
	return s.courseRepo
}

func (s *storagePg) Vacancy() repo.VacancyStorageI {
	return s.vacancyRepo
}
func (s *storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}
