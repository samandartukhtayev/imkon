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

// @Router /vacancies [post]
// @Summary Create a vacancy
// @Description Create a vacancy
// @Tags vacancies
// @Accept json
// @Produce json
// @Param vacancy body models.CreateVacancyReq true "vacancy"
// @Success 201 {object} models.VacancyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateVacancy(c *gin.Context) {
	var (
		req models.CreateVacancyReq
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Vacancy().CreateVacancy(&repo.CreateVacancyReq{
		Name:       req.Name,
		Address:    req.Address,
		JobType:    req.JobType,
		Info:       req.Info,
		MinSalary:  req.MinSalary,
		MaxSalary:  req.MaxSalary,
		BusinessId: req.BusinessId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.VacancyResponse{
		Id:   resp.Id,
		Name: resp.Name,
		CategoryInfo: models.Category{
			Id:   resp.CategoryInfo.Id,
			Name: resp.CategoryInfo.Name,
		},
		ImageUrl:  resp.ImageUrl,
		Address:   resp.Address,
		JobType:   resp.JobType,
		Info:      resp.Info,
		MinSalary: resp.MinSalary,
		MaxSalary: resp.MaxSalary,
		BusinessInfo: models.BusinessResponse{
			Id:               resp.BusinessInfo.Id,
			Name:             resp.BusinessInfo.Name,
			Address:          resp.BusinessInfo.Address,
			ImageUrl:         resp.BusinessInfo.ImageUrl,
			Info:             resp.BusinessInfo.Info,
			Email:            resp.BusinessInfo.Email,
			PhoneNumber:      resp.BusinessInfo.PhoneNumber,
			WebSite:          resp.BusinessInfo.WebSite,
			TelegramAccount:  resp.BusinessInfo.TelegramAccount,
			InstagramAccount: resp.BusinessInfo.InstagramAccount,
			LinkedInAccount:  resp.BusinessInfo.LinkedInAccount,
			CreatedAt:        resp.BusinessInfo.CreatedAt,
		},
		ViewsCount: resp.ViewsCount,
		CreatedAt:  resp.CreatedAt,
	})
}

// @Router /vacancies/{id} [get]
// @Summary Get vacancy by id
// @Description Get vacancy by id
// @Tags vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.VacancyResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vacancy, err := h.storage.Vacancy().GetVacancy(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.VacancyResponse{
		Id:   vacancy.Id,
		Name: vacancy.Name,
		CategoryInfo: models.Category{
			Id:   vacancy.CategoryInfo.Id,
			Name: vacancy.CategoryInfo.Name,
		},
		ImageUrl:  vacancy.ImageUrl,
		Address:   vacancy.Address,
		JobType:   vacancy.JobType,
		Info:      vacancy.Info,
		MinSalary: vacancy.MinSalary,
		MaxSalary: vacancy.MaxSalary,
		BusinessInfo: models.BusinessResponse{
			Id:               vacancy.BusinessInfo.Id,
			Name:             vacancy.BusinessInfo.Name,
			Address:          vacancy.BusinessInfo.Address,
			ImageUrl:         vacancy.BusinessInfo.ImageUrl,
			Info:             vacancy.BusinessInfo.Info,
			Email:            vacancy.BusinessInfo.Email,
			PhoneNumber:      vacancy.BusinessInfo.PhoneNumber,
			WebSite:          vacancy.BusinessInfo.WebSite,
			TelegramAccount:  vacancy.BusinessInfo.TelegramAccount,
			InstagramAccount: vacancy.BusinessInfo.InstagramAccount,
			LinkedInAccount:  vacancy.BusinessInfo.LinkedInAccount,
			CreatedAt:        vacancy.BusinessInfo.CreatedAt,
		},
		ViewsCount: vacancy.ViewsCount,
		CreatedAt:  vacancy.CreatedAt,
	})
}

// @Router /vacancies [get]
// @Summary Get all vacancies
// @Description Get all vacancies
// @Tags vacancies
// @Accept json
// @Produce json
// @Param filter query models.GetAllVacanciesReq false "Filter"
// @Success 200 {object} models.GetAllVacanciesRes
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllVacancies(c *gin.Context) {
	req, err := validateGetAllvacanciesParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	result, err := h.storage.Vacancy().GetAllVacancies(&repo.GetAllVacanciesReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getvacanciesResponse(result))
}

func getvacanciesResponse(data *repo.GetAllVacanciesRes) *models.GetAllVacanciesRes {
	response := models.GetAllVacanciesRes{
		Vacancies: make([]*models.VacancyResponse, 0),
		Count:     data.Count,
	}

	for _, vacancy := range data.Vacancies {
		vacncy := &models.VacancyResponse{
			Id:   vacancy.Id,
			Name: vacancy.Name,
			CategoryInfo: models.Category{
				Id:   vacancy.CategoryInfo.Id,
				Name: vacancy.CategoryInfo.Name,
			},
			ImageUrl:  vacancy.ImageUrl,
			Address:   vacancy.Address,
			JobType:   vacancy.JobType,
			Info:      vacancy.Info,
			MinSalary: vacancy.MinSalary,
			MaxSalary: vacancy.MaxSalary,
			BusinessInfo: models.BusinessResponse{
				Id:               vacancy.BusinessInfo.Id,
				Name:             vacancy.BusinessInfo.Name,
				Address:          vacancy.BusinessInfo.Address,
				ImageUrl:         vacancy.BusinessInfo.ImageUrl,
				Info:             vacancy.BusinessInfo.Info,
				Email:            vacancy.BusinessInfo.Email,
				PhoneNumber:      vacancy.BusinessInfo.PhoneNumber,
				WebSite:          vacancy.BusinessInfo.WebSite,
				TelegramAccount:  vacancy.BusinessInfo.TelegramAccount,
				InstagramAccount: vacancy.BusinessInfo.InstagramAccount,
				LinkedInAccount:  vacancy.BusinessInfo.LinkedInAccount,
				CreatedAt:        vacancy.BusinessInfo.CreatedAt,
			},
			ViewsCount: vacancy.ViewsCount,
			CreatedAt:  vacancy.CreatedAt,
		}
		response.Vacancies = append(response.Vacancies, vacncy)
	}

	return &response
}

func validateGetAllvacanciesParams(c *gin.Context) (*models.GetAllVacanciesReq, error) {
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

	return &models.GetAllVacanciesReq{
		Limit:      int64(limit),
		Page:       int64(page),
		Search:     c.Query("search"),
		SortByDate: c.Query("sort_by_date"),
	}, nil
}

// @Router /vacancies/{id} [put]
// @Summary Update a vacancy
// @Description Update a vacancy
// @Tags vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param vacancy body models.UpdateVacancyReq true "vacancy"
// @Success 200 {object} models.VacancyResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateVacancy(ctx *gin.Context) {
	var req models.UpdateVacancyReq

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

	vacancy, err := h.storage.Vacancy().UpdateVacancy(&repo.UpdateVacancyReq{
		Id:        int64(Id),
		Name:      req.Name,
		JobType:   req.JobType,
		Info:      req.Info,
		MinSalary: req.MaxSalary,
		MaxSalary: req.MaxSalary,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.VacancyResponse{
		Id:   vacancy.Id,
		Name: vacancy.Name,
		CategoryInfo: models.Category{
			Id:   vacancy.CategoryInfo.Id,
			Name: vacancy.CategoryInfo.Name,
		},
		ImageUrl:  vacancy.ImageUrl,
		Address:   vacancy.Address,
		JobType:   vacancy.JobType,
		Info:      vacancy.Info,
		MinSalary: vacancy.MinSalary,
		MaxSalary: vacancy.MaxSalary,
		BusinessInfo: models.BusinessResponse{
			Id:               vacancy.BusinessInfo.Id,
			Name:             vacancy.BusinessInfo.Name,
			Address:          vacancy.BusinessInfo.Address,
			ImageUrl:         vacancy.BusinessInfo.ImageUrl,
			Info:             vacancy.BusinessInfo.Info,
			Email:            vacancy.BusinessInfo.Email,
			PhoneNumber:      vacancy.BusinessInfo.PhoneNumber,
			WebSite:          vacancy.BusinessInfo.WebSite,
			TelegramAccount:  vacancy.BusinessInfo.TelegramAccount,
			InstagramAccount: vacancy.BusinessInfo.InstagramAccount,
			LinkedInAccount:  vacancy.BusinessInfo.LinkedInAccount,
			CreatedAt:        vacancy.BusinessInfo.CreatedAt,
		},
		ViewsCount: vacancy.ViewsCount,
		CreatedAt:  vacancy.CreatedAt,
	})
}

// @Summary Delete a Vacancy
// @Description Delete a vacancy
// @Tags vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /vacancies/{id} [delete]
func (h *handlerV1) DeleteVacancy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Vacancy().DeleteVacancy(int64(id))
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

// @Router /vacancies/image-upload/{id} [post]
// @Summary File image upload
// @Description File image upload
// @Tags vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"                 //TODO: MiddleWare'dan olish kerak
// @Param file formData file true "File"
// @Success 200 {object} models.VacancyResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) VacanciesImageUpload(c *gin.Context) {
	var file File

	VacancyId, err := strconv.Atoi(c.Param("id"))
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

	vacancy, err := h.storage.Vacancy().UploadVacancyImage(&repo.UploadVacancyImageReq{
		VacancyId: int64(VacancyId),
		ImageUrl:  filePath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.VacancyResponse{
		Id:   vacancy.Id,
		Name: vacancy.Name,
		CategoryInfo: models.Category{
			Id:   vacancy.CategoryInfo.Id,
			Name: vacancy.CategoryInfo.Name,
		},
		ImageUrl:  vacancy.ImageUrl,
		Address:   vacancy.Address,
		JobType:   vacancy.JobType,
		Info:      vacancy.Info,
		MinSalary: vacancy.MinSalary,
		MaxSalary: vacancy.MaxSalary,
		BusinessInfo: models.BusinessResponse{
			Id:               vacancy.BusinessInfo.Id,
			Name:             vacancy.BusinessInfo.Name,
			Address:          vacancy.BusinessInfo.Address,
			ImageUrl:         vacancy.BusinessInfo.ImageUrl,
			Info:             vacancy.BusinessInfo.Info,
			Email:            vacancy.BusinessInfo.Email,
			PhoneNumber:      vacancy.BusinessInfo.PhoneNumber,
			WebSite:          vacancy.BusinessInfo.WebSite,
			TelegramAccount:  vacancy.BusinessInfo.TelegramAccount,
			InstagramAccount: vacancy.BusinessInfo.InstagramAccount,
			LinkedInAccount:  vacancy.BusinessInfo.LinkedInAccount,
			CreatedAt:        vacancy.BusinessInfo.CreatedAt,
		},
		ViewsCount: vacancy.ViewsCount,
		CreatedAt:  vacancy.CreatedAt,
	})
}
