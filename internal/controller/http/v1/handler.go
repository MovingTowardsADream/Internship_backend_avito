package v1

import (
	"Internship_backend_avito/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			accounts := v1.Group("/accounts")
			{
				accounts.POST("/create", h.AccountCreate)
				accounts.POST("/deposit", h.AccountDeposit)
				accounts.POST("/withdraw", h.AccountWithDraw)
				accounts.POST("/transfer", h.AccountTransfer)
				accounts.GET("/:id", h.AccountsGet)
			}
			operations := v1.Group("/operations")
			{
				operations.GET("/history", h.OperationsHistory)
				operations.GET("/link", h.OperationsLink)
				operations.GET("/file", h.OperationsFile)
			}
			product := v1.Group("/product")
			{
				product.POST("/create", h.ProductCreate)
				product.GET("/:id", h.ProductGet)
			}
			reservation := v1.Group("/reservation", h.userIdentity)
			{
				reservation.POST("/create", h.ReservationCreate)
				reservation.POST("/revenue", h.ReservationRevenue)
				reservation.POST("/refund", h.ReservationRefund)
			}
		}
	}

	return router
}
