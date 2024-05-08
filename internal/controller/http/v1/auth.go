package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
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

	token, err := h.service.Authorization.GenerateToken(service.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
