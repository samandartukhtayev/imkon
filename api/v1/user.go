package v1

import (
	"mime/multipart"
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

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserReq true "user"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req models.CreateUserReq
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

	resp, err := h.storage.User().CreateUser(&repo.CreateUserReq{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.UserResponse{
		Id:           resp.Id,
		FirstName:    resp.FirstName,
		LastName:     resp.LastName,
		Email:        resp.Email,
		PhoneNumber:  resp.PhoneNumber,
		ImageUrl:     resp.ImageUrl,
		PortfoliaUrl: resp.PortfoliaUrl,
		CreatedAt:    resp.CreatedAt,
	})
}

// @Router /users/{id} [get]
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.UserResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := h.storage.User().GetUser(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ImageUrl:     user.ImageUrl,
		PortfoliaUrl: user.PortfoliaUrl,
		CreatedAt:    user.CreatedAt,
	})
}

// @Router /users [get]
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Param filter query models.GetAllUsersReq false "Filter"
// @Success 200 {object} models.GetAllUsersResp
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	result, err := h.storage.User().GetAllUsers(&repo.GetAllUsersReq{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getUsersResponse(result))
}

func getUsersResponse(data *repo.GetAllUsersResp) *models.GetAllUsersResp {
	response := models.GetAllUsersResp{
		Users: make([]*models.UserResponse, 0),
		Count: data.Count,
	}

	for _, user := range data.Users {
		usr := &models.UserResponse{
			Id:           user.Id,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			ImageUrl:     user.ImageUrl,
			PortfoliaUrl: user.PortfoliaUrl,
			CreatedAt:    user.CreatedAt,
		}
		response.Users = append(response.Users, usr)
	}

	return &response
}

func validateGetAllParams(c *gin.Context) (*models.GetAllUsersReq, error) {
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

	return &models.GetAllUsersReq{
		Limit:      int64(limit),
		Page:       int64(page),
		Search:     c.Query("search"),
		SortByDate: c.Query("sort_by_date"),
	}, nil
}

// @Router /users-accepted [get]
// @Summary Get all accepted users
// @Description Get all accepted users
// @Tags users
// @Accept json
// @Produce json
// @Param filter query models.GetAllAcceptedUsersReq false "Filter"
// @Success 200 {object} models.GetAllUsersResp
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllAcceptedUsers(c *gin.Context) {
	req, err := validateGetAcceptedUserAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	result, err := h.storage.User().GetAllAcceptedUsers(&repo.GetAllAcceptedUsersReq{
		Page:       req.Page,
		Limit:      req.Limit,
		VacancyId:  req.VacancyId,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, getUsersResponse(result))
}

func validateGetAcceptedUserAllParams(c *gin.Context) (*models.GetAllAcceptedUsersReq, error) {
	var (
		limit     int = 10
		page      int = 1
		vacancyId int
		err       error
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
	if c.Query("vacancy_id") != "" {
		vacancyId, err = strconv.Atoi(c.Query("vacancy_id"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllAcceptedUsersReq{
		Limit:      int64(limit),
		Page:       int64(page),
		Search:     c.Query("search"),
		VacancyId:  int64(vacancyId),
		SortByDate: c.Query("sort_by_date"),
	}, nil
}

// @Router /users/{id} [put]
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param user body models.UpdateUserReq true "user"
// @Success 200 {object} models.UserResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var req models.UpdateUserReq

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

	user, err := h.storage.User().UpdateUser(&repo.UpdateUserReq{
		Id:          int64(Id),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.UserResponse{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ImageUrl:     user.ImageUrl,
		PortfoliaUrl: user.PortfoliaUrl,
		CreatedAt:    user.CreatedAt,
	})
}

// @Summary Delete a User
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (h *handlerV1) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.storage.User().DeleteUser(int64(id))
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

// @Router /users/image-upload/{id} [post]
// @Summary File image upload
// @Description File image upload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"                 //TODO: MiddleWare'dan olish kerak
// @Param file formData file true "File"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UsersImageUpload(c *gin.Context) {
	var file File

	UserId, err := strconv.Atoi(c.Param("id"))
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

	user, err := h.storage.User().UploadUserImage(&repo.UploadUserImageReq{
		UserId:   int64(UserId),
		ImageUrl: filePath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ImageUrl:     user.ImageUrl,
		PortfoliaUrl: user.PortfoliaUrl,
		CreatedAt:    user.CreatedAt,
	})
}

// @Router /users/portfolia-upload/{id} [post]
// @Summary File portfolia upload
// @Description File portfolia upload
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID"                 //TODO: MiddleWare'dan olish kerak
// @Param file formData file true "File"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UsersPortfoliaUpload(c *gin.Context) {
	var file File

	UserId, err := strconv.Atoi(c.Param("id"))
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

	user, err := h.storage.User().UploadUserPortfolia(&repo.UploadUserPortfoliaReq{
		UserId:       int64(UserId),
		PortfoliaUrl: filePath,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ImageUrl:     user.ImageUrl,
		PortfoliaUrl: user.PortfoliaUrl,
		CreatedAt:    user.CreatedAt,
	})
}
