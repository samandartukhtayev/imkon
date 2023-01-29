package repo

import "time"

type Course struct {
	Id          int64
	Name        string
	CoursePrice int
	CategoryId  int64
	Info        string
	BusinessId  int64
	ImageUrl    string
	SaleOf      int8
	CreatedAt   time.Time
}

type CourseResponse struct {
	Id          int64
	Name        string
	CoursePrice int
	Category    Category
	Info        string
	Business    BusinessResponse
	ImageUrl    string
	SaleOf      int8
	CreatedAt   time.Time
}

type CreateCourseReq struct {
	Name        string
	CoursePrice int
	Info        string
	BusinessId  int64
	SaleOf      int8
}

type UpdateCourseReq struct {
	Id          int64
	Name        string
	CoursePrice int
	Info        string
	SaleOf      int8
}

type UploadCourseImageReq struct {
	CourseId int64
	ImageUrl string
}

type GetAllCoursesReq struct {
	Page       int64
	Limit      int64
	Search     string
	SortByDate string
}

type GetAllCoursesRes struct {
	Courses []*CourseResponse
	Count   int64
}

type CourseStorageI interface {
	CreateCourse(*CreateCourseReq) (*CourseResponse, error)
	GetCourse(int64) (*CourseResponse, error)
	GetAllCourses(*GetAllCoursesReq) (*GetAllCoursesRes, error)
	UpdateCourse(*UpdateCourseReq) (*CourseResponse, error)
	DeleteCourse(int64) error
	UploadCourseImage(*UploadCourseImageReq) (*CourseResponse, error)
}
