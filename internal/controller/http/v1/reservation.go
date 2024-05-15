package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reservationCreateInput struct {
	AccountId int `json:"account_id" binding:"required"`
	ProductId int `json:"product_id" binding:"required"`
	OrderId   int `json:"order_id" binding:"required"`
	Amount    int `json:"amount" binding:"required"`
}

func (h *Handler) ReservationCreate(c *gin.Context) {
	var input reservationCreateInput

	if err := c.Bind(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.service.Reservation.CreateReservation(c.Request.Context(), service.CreateReservationInput{
		AccountId: input.AccountId,
		ProductId: input.ProductId,
		OrderId:   input.OrderId,
		Amount:    input.Amount,
	})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Id int `json:"id" binding:"required"`
	}

	c.JSON(http.StatusOK, response{
		Id: id,
	})
}

func (h *Handler) ReservationRevenue(c *gin.Context) {

}

func (h *Handler) ReservationRefund(c *gin.Context) {

}
