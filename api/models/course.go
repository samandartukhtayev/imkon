package models

import "time"

type Course struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
 	CoursePrice int       `json:"course_price"`
	CategoryId  int64     `json:"category_id"`
	Info        string    `json:"info"`
	BusinessId  int64     `json:"business_id"`
	ImageUrl    string    `json:"image_url"`
	SaleOf      int8      `json:"sale_of"`
	CreatedAt   time.Time `json:"created_at"`
}

type CourseResponse struct {
	Id          int64            `json:"id"`
	Name        string           `json:"name"`
 	CoursePrice int              `json:"course_price"`
	Category    Category         `json:"category"`
	Info        string           `json:"info"`
	Business    BusinessResponse `json:"business"`
	ImageUrl    string           `json:"image_url"`
	SaleOf      int8             `json:"sale_of"`
	CreatedAt   time.Time        `json:"created_at"`
}

type CreateCourseReq struct {
	Name        string `json:"name"`
 	CoursePrice int    `json:"course_price"`
	Info        string `json:"info"`
	BusinessId  int64  `json:"business_id"`
	SaleOf      int8   `json:"sale_of"`
}

type UpdateCourseReq struct {
	Name        string `json:"name"`
 	CoursePrice int    `json:"course_price"`
	Info        string `json:"info"`
	SaleOf      int8   `json:"sale_of"`
}

type UploadCourseImageReq struct {
	CourseId int64  `json:"course_id"`
	ImageUrl string `json:"image_url"`
}

type GetAllCoursesReq struct {
	Page       int64  `json:"page" binding:"required" default:"1"`
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllCoursesRes struct {
	Courses []*CourseResponse `json:"courses"`
	Count   int64             `json:"count"`
}
