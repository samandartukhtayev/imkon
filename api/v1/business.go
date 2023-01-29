package v1

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samandartukhtayev/imkon/api/models"
	"github.com/samandartukhtayev/imkon/pkg/utils"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

// @Router /businesses [post]
// @Summary Create a business
// @Description Create a business
// @Tags businesses
// @Accept json
// @Produce json
// @Param business body models.CreateBusinessReq true "business"
// @Success 201 {object} models.BusinessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateBusiness(c *gin.Context) {
	var (
		req models.CreateBusinessReq
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Business().CreateBusiness(&repo.CreateBusinessReq{
		Name:             req.Name,
		Password:         hashedPassword,
		Address:          req.Address,
		Info:             req.Info,
		Email:            req.Email,
		PhoneNumber:      req.PhoneNumber,
		WebSite:          req.WebSite,
		TelegramAccount:  req.TelegramAccount,
		InstagramAccount: req.InstagramAccount,
		LinkedInAccount:  req.LinkedInAccount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.BusinessResponse{
		Id:               resp.Id,
		Name:             resp.Name,
		Address:          resp.Address,
		ImageUrl:         resp.ImageUrl,
		Info:             resp.Info,
		Email:            resp.Email,
		PhoneNumber:      resp.PhoneNumber,
		WebSite:          resp.WebSite,
		TelegramAccount:  resp.TelegramAccount,
		InstagramAccount: resp.InstagramAccount,
		LinkedInAccount:  resp.LinkedInAccount,
		CreatedAt:        resp.CreatedAt,
	})
}

// @Router /businesses/{id} [get]
// @Summary Get business by id
// @Description Get business by id
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.BusinessResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetBusiness(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	business, err := h.storage.Business().GetBusiness(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.BusinessResponse{
		Id:               business.Id,
		Name:             business.Name,
		Address:          business.Address,
		ImageUrl:         business.ImageUrl,
		Info:             business.Info,
		Email:            business.Email,
		PhoneNumber:      business.PhoneNumber,
		WebSite:          business.WebSite,
		TelegramAccount:  business.TelegramAccount,
		InstagramAccount: business.InstagramAccount,
		LinkedInAccount:  business.LinkedInAccount,
		CreatedAt:        business.CreatedAt,
	})
}

// @Router /businesses/{id} [put]
// @Summary Update a business
// @Description Update a business
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param business body models.UpdateBusinessReq true "business"
// @Success 200 {object} models.BusinessResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateBusiness(ctx *gin.Context) {
	var req models.UpdateBusinessReq

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

	business, err := h.storage.Business().UpdateBusiness(&repo.UpdateBusinessReq{
		Id:               int64(Id),
		Name:             req.Name,
		Address:          req.Address,
		Info:             req.Info,
		PhoneNumber:      req.PhoneNumber,
		WebSite:          req.WebSite,
		InstagramAccount: req.InstagramAccount,
		TelegramAccount:  req.TelegramAccount,
		LinkedInAccount:  req.LinkedInAccount,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.BusinessResponse{
		Id:               business.Id,
		Name:             business.Name,
		Address:          business.Address,
		ImageUrl:         business.ImageUrl,
		Info:             business.Info,
		Email:            business.Email,
		PhoneNumber:      business.PhoneNumber,
		WebSite:          business.WebSite,
		TelegramAccount:  business.TelegramAccount,
		InstagramAccount: business.InstagramAccount,
		LinkedInAccount:  business.LinkedInAccount,
		CreatedAt:        business.CreatedAt,
	})
}

// @Router /businesses [get]
// @Summary Get all businesses
// @Description Get all businesses
// @Tags businesses
// @Accept json
// @Produce json
// @Param filter query models.GetAllBusinessesReq false "Filter"
// @Success 200 {object} models.GetAllBusinessesResp
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllBusinesses(c *gin.Context) {
	req, err := validateGetAllBusinessesParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	result, err := h.storage.Business().GetAllBusinesses(&repo.GetAllBusinessesReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getBusinessesResponse(result))
}

func getBusinessesResponse(data *repo.GetAllBusinessesRes) *models.GetAllBusinessesResp {
	response := models.GetAllBusinessesResp{
		Businesses: make([]*models.BusinessResponse, 0),
		Count:      data.Count,
	}

	for _, business := range data.Businesses {
		bsness := &models.BusinessResponse{
			Id:               business.Id,
			Name:             business.Name,
			Address:          business.Address,
			ImageUrl:         business.ImageUrl,
			Info:             business.Info,
			Email:            business.Email,
			PhoneNumber:      business.PhoneNumber,
			WebSite:          business.WebSite,
			TelegramAccount:  business.TelegramAccount,
			InstagramAccount: business.InstagramAccount,
			LinkedInAccount:  business.LinkedInAccount,
			CreatedAt:        business.CreatedAt,
		}
		response.Businesses = append(response.Businesses, bsness)
	}

	return &response
}

func validateGetAllBusinessesParams(c *gin.Context) (*models.GetAllBusinessesReq, error) {
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

	return &models.GetAllBusinessesReq{
		Limit:      int64(limit),
		Page:       int64(page),
		Search:     c.Query("search"),
		SortByDate: c.Query("sort_by_date"),
	}, nil
}

// @Summary Delete a Business
// @Description Delete a business
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /businesses/{id} [delete]
func (h *handlerV1) DeleteBusiness(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.Business().DeleteBusiness(int64(id))
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

// @Router /businesses/image-upload/{id} [post]
// @Summary File image upload
// @Description File image upload
// @Tags businesses
// @Accept json
// @Produce json
// @Param id path int true "ID"                 //TODO: MiddleWare'dan olish kerak
// @Param file formData file true "File"
// @Success 200 {object} models.BusinessResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) BusinessesImageUpload(c *gin.Context) {
	var file File

	BusinessId, err := strconv.Atoi(c.Param("id"))
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

	business, err := h.storage.Business().UploadBusinessImage(&repo.UploadBusinessImageReq{
		BusinessId: int64(BusinessId),
		ImageUrl:   filePath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.BusinessResponse{
		Id:               business.Id,
		Name:             business.Name,
		Address:          business.Address,
		ImageUrl:         business.ImageUrl,
		Info:             business.Info,
		Email:            business.Email,
		PhoneNumber:      business.PhoneNumber,
		WebSite:          business.WebSite,
		TelegramAccount:  business.TelegramAccount,
		InstagramAccount: business.InstagramAccount,
		LinkedInAccount:  business.LinkedInAccount,
		CreatedAt:        business.CreatedAt,
	})
}
