package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	limitDefault  = 1
	offsetDefault = 0
)

type getHistoryInput struct {
	AccountId int    `json:"account_id" binding:"required"`
	SortType  string `json:"sort_type" binding:"required"`
}

func (h *Handler) OperationsHistory(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = limitDefault
	}

	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = offsetDefault
	}

	var input getHistoryInput

	if err = c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	operations, err := h.service.Operation.OperationsHistory(c.Request.Context(), service.OperationHistoryInput{
		AccountId: input.AccountId,
		SortType:  input.SortType,
		Offset:    offset,
		Limit:     limit,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Operations []service.OperationHistoryOutput `json:"operations"`
	}

	c.JSON(http.StatusOK, response{
		Operations: operations,
	})
}

func (h *Handler) OperationsLink(c *gin.Context) {

}

func (h *Handler) OperationsFile(c *gin.Context) {

}
