package v1

import (
	"Internship_backend_avito/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productCreateInput struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) ProductCreate(c *gin.Context) {
	var input productCreateInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.service.Product.CreateProduct(c.Request.Context(), input.Name)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Id int `json:"id"`
	}

	c.JSON(http.StatusCreated, response{
		Id: id,
	})
}

func (h *Handler) ProductGet(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	product, err := h.service.Product.GetProductById(c.Request.Context(), productId)
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Product entity.Product `json:"product"`
	}

	c.JSON(http.StatusOK, response{
		Product: product,
	})
}
