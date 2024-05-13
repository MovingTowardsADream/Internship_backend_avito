package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(c.Request.Context(), service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Id int `json:"id"`
	}

	c.JSON(http.StatusCreated, response{
		Id: id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(c.Request.Context(), service.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	c.JSON(http.StatusOK, response{
		Token: token,
	})
}
