package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) AccountCreate(c *gin.Context) {
	id, err := h.service.Account.CreateAccount(c.Request.Context())

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Id int `json:"id"`
	}

	c.JSON(http.StatusOK, response{
		Id: id,
	})
}

type accountDepositInput struct {
	Id     int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

func (h *Handler) AccountDeposit(c *gin.Context) {
	var input accountDepositInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.Account.AccountDeposit(c.Request.Context(), service.AccountDepositInput{
		Id:     input.Id,
		Amount: input.Amount,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
	})
}

type accountWithdrawInput struct {
	Id     int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

func (h *Handler) AccountWithDraw(c *gin.Context) {
	var input accountWithdrawInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.Account.Withdraw(c.Request.Context(), service.AccountWithdrawInput{
		Id:     input.Id,
		Amount: input.Amount,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
	})
}

type accountTransferInput struct {
	IdFrom int `json:"id_from" binding:"required"`
	IdTo   int `json:"id_to" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

func (h *Handler) AccountTransfer(c *gin.Context) {
	var input accountTransferInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.Account.Transfer(c.Request.Context(), service.AccountTransferInput{
		IdFrom: input.IdFrom,
		IdTo:   input.IdTo,
		Amount: input.Amount,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success",
	})
}

func (h *Handler) AccountsGet(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	account, err := h.service.Account.GetAccountById(c.Request.Context(), accountId)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Id      int `json:"id" binding:"required"`
		Balance int `json:"balance" binding:"required"`
	}

	c.JSON(http.StatusOK, response{
		Id:      account.Id,
		Balance: account.Balance,
	})
}
