package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samandartukhtayev/imkon/api/models"
	"github.com/samandartukhtayev/imkon/storage/repo"
)

// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryReq true "category"
// @Success 201 {object} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		req models.CreateCategoryReq
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().CreateCategory(&repo.CreateCategoryReq{
		Name: req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Category{
		Id:   resp.Id,
		Name: resp.Name,
	})
}
