package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandartukhtayev/imkon/api/models"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

// @Router /accepted-vacancies [post]
// @Summary Create a accepted_vacancy
// @Description Create a accepted_vacancy
// @Tags accepted-vacancies
// @Accept json
// @Produce json
// @Param accepted_vacancy body models.AcceptVacancyReq true "accepted_vacancy"
// @Success 200 {object} models.AcceptVacancyRes
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) AcceptVacancy(c *gin.Context) {
	var (
		req models.AcceptVacancyReq
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.AcceptVacancy().AcceptVacancy(&repo.AcceptVacancyReq{
		UserId:    req.UserId,
		VacancyId: req.VacancyId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AcceptVacancyRes{
		Id: resp.Id,
		UserInfo: models.UserResponse{
			Id:           resp.UserInfo.Id,
			FirstName:    resp.UserInfo.FirstName,
			LastName:     resp.UserInfo.LastName,
			Email:        resp.UserInfo.Email,
			PhoneNumber:  resp.UserInfo.PhoneNumber,
			ImageUrl:     resp.UserInfo.ImageUrl,
			PortfoliaUrl: resp.UserInfo.PortfoliaUrl,
			CreatedAt:    resp.UserInfo.CreatedAt,
		},
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
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /accepted-vacancies-by-user/{id} [get]
// @Summary Get vacancy by user id
// @Description Get vacancy by user id
// @Tags accepted-vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.AcceptedVacanciesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAcceptedVacanciesByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	accepteds, err := h.storage.AcceptVacancy().GetAcceptedVacanciesByUserId(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var result models.AcceptedVacanciesResponse
	for _, accepted := range accepteds.AcceptedVacancies {
		result.AcceptedVacancies = append(result.AcceptedVacancies, &models.AcceptVacancyRes{
			Id: accepted.Id,
			UserInfo: models.UserResponse{
				Id:           accepted.UserInfo.Id,
				FirstName:    accepted.UserInfo.FirstName,
				LastName:     accepted.UserInfo.LastName,
				Email:        accepted.UserInfo.Email,
				PhoneNumber:  accepted.UserInfo.PhoneNumber,
				ImageUrl:     accepted.UserInfo.ImageUrl,
				PortfoliaUrl: accepted.UserInfo.PortfoliaUrl,
				CreatedAt:    accepted.UserInfo.CreatedAt,
			},
			BusinessInfo: models.BusinessResponse{
				Id:               accepted.BusinessInfo.Id,
				Name:             accepted.BusinessInfo.Name,
				Address:          accepted.BusinessInfo.Address,
				ImageUrl:         accepted.BusinessInfo.ImageUrl,
				Info:             accepted.BusinessInfo.Info,
				Email:            accepted.BusinessInfo.Email,
				PhoneNumber:      accepted.BusinessInfo.PhoneNumber,
				WebSite:          accepted.BusinessInfo.WebSite,
				TelegramAccount:  accepted.BusinessInfo.TelegramAccount,
				InstagramAccount: accepted.BusinessInfo.InstagramAccount,
				LinkedInAccount:  accepted.BusinessInfo.LinkedInAccount,
				CreatedAt:        accepted.BusinessInfo.CreatedAt,
			},
			CreatedAt: accepted.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, models.AcceptedVacanciesResponse{
		AcceptedVacancies: result.AcceptedVacancies,
		Count:             accepteds.Count,
	})
}

// @Router /accepted-vacancies-by-business/{id} [get]
// @Summary Get vacancy by business id
// @Description Get vacancy by business id
// @Tags accepted-vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.AcceptedVacanciesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAcceptedVacanciesByBusinessId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	accepteds, err := h.storage.AcceptVacancy().GetAcceptedVacanciesByBusinessId(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var result models.AcceptedVacanciesResponse
	for _, accepted := range accepteds.AcceptedVacancies {
		result.AcceptedVacancies = append(result.AcceptedVacancies, &models.AcceptVacancyRes{
			Id: accepted.Id,
			UserInfo: models.UserResponse{
				Id:           accepted.UserInfo.Id,
				FirstName:    accepted.UserInfo.FirstName,
				LastName:     accepted.UserInfo.LastName,
				Email:        accepted.UserInfo.Email,
				PhoneNumber:  accepted.UserInfo.PhoneNumber,
				ImageUrl:     accepted.UserInfo.ImageUrl,
				PortfoliaUrl: accepted.UserInfo.PortfoliaUrl,
				CreatedAt:    accepted.UserInfo.CreatedAt,
			},
			BusinessInfo: models.BusinessResponse{
				Id:               accepted.BusinessInfo.Id,
				Name:             accepted.BusinessInfo.Name,
				Address:          accepted.BusinessInfo.Address,
				ImageUrl:         accepted.BusinessInfo.ImageUrl,
				Info:             accepted.BusinessInfo.Info,
				Email:            accepted.BusinessInfo.Email,
				PhoneNumber:      accepted.BusinessInfo.PhoneNumber,
				WebSite:          accepted.BusinessInfo.WebSite,
				TelegramAccount:  accepted.BusinessInfo.TelegramAccount,
				InstagramAccount: accepted.BusinessInfo.InstagramAccount,
				LinkedInAccount:  accepted.BusinessInfo.LinkedInAccount,
				CreatedAt:        accepted.BusinessInfo.CreatedAt,
			},
			CreatedAt: accepted.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, models.AcceptedVacanciesResponse{
		AcceptedVacancies: result.AcceptedVacancies,
		Count:             accepteds.Count,
	})
}

// @Router /accepted-vacancies-by-id/{id} [get]
// @Summary Get vacancy by id
// @Description Get vacancy by id
// @Tags accepted-vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.AcceptVacancyRes
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAcceptedVacanciesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.AcceptVacancy().GetAcceptedVacanciesById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.AcceptVacancyRes{
		Id: resp.Id,
		UserInfo: models.UserResponse{
			Id:           resp.UserInfo.Id,
			FirstName:    resp.UserInfo.FirstName,
			LastName:     resp.UserInfo.LastName,
			Email:        resp.UserInfo.Email,
			PhoneNumber:  resp.UserInfo.PhoneNumber,
			ImageUrl:     resp.UserInfo.ImageUrl,
			PortfoliaUrl: resp.UserInfo.PortfoliaUrl,
			CreatedAt:    resp.UserInfo.CreatedAt,
		},
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
		CreatedAt: resp.CreatedAt,
	})
}

// @Summary Delete a accepted_vacancy
// @Description Delete a accepted_vacancy
// @Tags accepted-vacancies
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /accepted-vacancy/{id} [delete]
func (h *handlerV1) DeleteAcceptedVacancy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.AcceptVacancy().DeleteAcceptedVacancy(int64(id))
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
