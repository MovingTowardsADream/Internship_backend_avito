package v1

import (
	"Internship_backend_avito/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

} // TODO del this

type getOperationsFileInput struct {
	Month int `json:"month" validate:"required"`
	Year  int `json:"year" validate:"required"`
}

func (h *Handler) OperationsFile(c *gin.Context) {
	var input getOperationsFileInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	file, err := h.service.Operation.OperationsFile(c.Request.Context(), input.Month, input.Year)
	if err != nil {
		logrus.Debugf("error while getting report file: %s", err.Error())
		NewErrorMessageResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	c.Header("Content-Disposition", "attachment; filename=operations.csv")
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Length", fmt.Sprintf("%d", len(file)))

	// Отправляем файл в ответ
	c.Data(http.StatusOK, "text/csv", file)

}
