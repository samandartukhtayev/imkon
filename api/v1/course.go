package v1

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samandartukhtayev/imkon/api/models"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

// @Router /courses [post]
// @Summary Create a course
// @Description Create a course
// @Tags courses
// @Accept json
// @Produce json
// @Param course body models.CreateCourseReq true "course"
// @Success 201 {object} models.CourseResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCourse(c *gin.Context) {
	var (
		req models.CreateCourseReq
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Course().CreateCourse(&repo.CreateCourseReq{
		Name:        req.Name,
		CoursePrice: req.CoursePrice,
		Info:        req.Info,
		BusinessId:  req.BusinessId,
		SaleOf:      req.SaleOf,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CourseResponse{
		Id:          resp.Id,
		Name:        resp.Name,
		CoursePrice: resp.CoursePrice,
		Category: models.Category{
			Id:   resp.Category.Id,
			Name: resp.Category.Name,
		},
		Info: resp.Info,
		Business: models.BusinessResponse{
			Id:               resp.Business.Id,
			Name:             resp.Business.Name,
			Address:          resp.Business.Address,
			ImageUrl:         resp.Business.ImageUrl,
			Info:             resp.Business.Info,
			Email:            resp.Business.Email,
			PhoneNumber:      resp.Business.PhoneNumber,
			WebSite:          resp.Business.WebSite,
			TelegramAccount:  resp.Business.TelegramAccount,
			InstagramAccount: resp.Business.InstagramAccount,
			LinkedInAccount:  resp.Business.LinkedInAccount,
			CreatedAt:        resp.Business.CreatedAt,
		},
		ImageUrl:  resp.ImageUrl,
		SaleOf:    resp.SaleOf,
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /courses/{id} [get]
// @Summary Get course by id
// @Description Get course by id
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.CourseResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	course, err := h.storage.Course().GetCourse(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.CourseResponse{
		Id:          course.Id,
		Name:        course.Name,
		CoursePrice: course.CoursePrice,
		Category: models.Category{
			Id:   course.Category.Id,
			Name: course.Category.Name,
		},
		Info: course.Info,
		Business: models.BusinessResponse{
			Id:               course.Business.Id,
			Name:             course.Business.Name,
			Address:          course.Business.Address,
			ImageUrl:         course.Business.ImageUrl,
			Info:             course.Business.Info,
			Email:            course.Business.Email,
			PhoneNumber:      course.Business.PhoneNumber,
			WebSite:          course.Business.WebSite,
			TelegramAccount:  course.Business.TelegramAccount,
			InstagramAccount: course.Business.InstagramAccount,
			LinkedInAccount:  course.Business.LinkedInAccount,
			CreatedAt:        course.Business.CreatedAt,
		},
		ImageUrl:  course.ImageUrl,
		SaleOf:    course.SaleOf,
		CreatedAt: course.CreatedAt,
	})
}

// @Router /courses [get]
// @Summary Get all courses
// @Description Get all courses
// @Tags courses
// @Accept json
// @Produce json
// @Param filter query models.GetAllCoursesReq false "Filter"
// @Success 200 {object} models.GetAllCoursesRes
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllCourses(c *gin.Context) {
	req, err := validateGetAllCoursesParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	result, err := h.storage.Course().GetAllCourses(&repo.GetAllCoursesReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getCoursesResponse(result))
}

func getCoursesResponse(data *repo.GetAllCoursesRes) *models.GetAllCoursesRes {
	response := models.GetAllCoursesRes{
		Courses: make([]*models.CourseResponse, 0),
		Count:   data.Count,
	}

	for _, course := range data.Courses {
		curse := &models.CourseResponse{
			Id:          course.Id,
			Name:        course.Name,
			CoursePrice: course.CoursePrice,
			Category: models.Category{
				Id:   course.Category.Id,
				Name: course.Category.Name,
			},
			Info: course.Info,
			Business: models.BusinessResponse{
				Id:               course.Business.Id,
				Name:             course.Business.Name,
				Address:          course.Business.Address,
				ImageUrl:         course.Business.ImageUrl,
				Info:             course.Business.Info,
				Email:            course.Business.Email,
				PhoneNumber:      course.Business.PhoneNumber,
				WebSite:          course.Business.WebSite,
				TelegramAccount:  course.Business.TelegramAccount,
				InstagramAccount: course.Business.InstagramAccount,
				LinkedInAccount:  course.Business.LinkedInAccount,
				CreatedAt:        course.Business.CreatedAt,
			},
			ImageUrl:  course.ImageUrl,
			SaleOf:    course.SaleOf,
			CreatedAt: course.CreatedAt,
		}
		response.Courses = append(response.Courses, curse)
	}

	return &response
}

func validateGetAllCoursesParams(c *gin.Context) (*models.GetAllCoursesReq, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllCoursesReq{
		Limit:      int64(limit),
		Page:       int64(page),
		Search:     c.Query("search"),
		SortByDate: c.Query("sort_by_date"),
	}, nil
}

// @Router /courses/{id} [put]
// @Summary Update a course
// @Description Update a course
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param course body models.UpdateCourseReq true "course"
// @Success 200 {object} models.CourseResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateCourse(ctx *gin.Context) {
	var req models.UpdateCourseReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	Id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	course, err := h.storage.Course().UpdateCourse(&repo.UpdateCourseReq{
		Id:          int64(Id),
		Name:        req.Name,
		CoursePrice: req.CoursePrice,
		Info:        req.Info,
		SaleOf:      req.SaleOf,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CourseResponse{
		Id:          course.Id,
		Name:        course.Name,
		CoursePrice: course.CoursePrice,
		Category: models.Category{
			Id:   course.Category.Id,
			Name: course.Category.Name,
		},
		Info: course.Info,
		Business: models.BusinessResponse{
			Id:               course.Business.Id,
			Name:             course.Business.Name,
			Address:          course.Business.Address,
			ImageUrl:         course.Business.ImageUrl,
			Info:             course.Business.Info,
			Email:            course.Business.Email,
			PhoneNumber:      course.Business.PhoneNumber,
			WebSite:          course.Business.WebSite,
			TelegramAccount:  course.Business.TelegramAccount,
			InstagramAccount: course.Business.InstagramAccount,
			LinkedInAccount:  course.Business.LinkedInAccount,
			CreatedAt:        course.Business.CreatedAt,
		},
		ImageUrl:  course.ImageUrl,
		SaleOf:    course.SaleOf,
		CreatedAt: course.CreatedAt,
	})
}

// @Summary Delete a Course
// @Description Delete a course
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /courses/{id} [delete]
func (h *handlerV1) DeleteCourse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Course().DeleteCourse(int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful delete method",
	})
}

// @Router /courses/image-upload/{id} [post]
// @Summary File image upload
// @Description File image upload
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "ID"                 //TODO: MiddleWare'dan olish kerak
// @Param file formData file true "File"
// @Success 200 {object} models.CourseResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CoursesImageUpload(c *gin.Context) {
	var file File

	CourseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})

		return
	}

	err = c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()
	if _, err := os.Stat(dst + "/media"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media", os.ModePerm)
	}

	filePath := "/media/" + fileName
	if err = c.SaveUploadedFile(file.File, dst+filePath); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	course, err := h.storage.Course().UploadCourseImage(&repo.UploadCourseImageReq{
		CourseId: int64(CourseId),
		ImageUrl: filePath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.CourseResponse{
		Id:          course.Id,
		Name:        course.Name,
		CoursePrice: course.CoursePrice,
		Category: models.Category{
			Id:   course.Category.Id,
			Name: course.Category.Name,
		},
		Info: course.Info,
		Business: models.BusinessResponse{
			Id:               course.Business.Id,
			Name:             course.Business.Name,
			Address:          course.Business.Address,
			ImageUrl:         course.Business.ImageUrl,
			Info:             course.Business.Info,
			Email:            course.Business.Email,
			PhoneNumber:      course.Business.PhoneNumber,
			WebSite:          course.Business.WebSite,
			TelegramAccount:  course.Business.TelegramAccount,
			InstagramAccount: course.Business.InstagramAccount,
			LinkedInAccount:  course.Business.LinkedInAccount,
			CreatedAt:        course.Business.CreatedAt,
		},
		ImageUrl:  course.ImageUrl,
		SaleOf:    course.SaleOf,
		CreatedAt: course.CreatedAt,
	})
}
