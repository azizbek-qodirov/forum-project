package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/docs"
	"api-gateway/api/handlers"
	"api-gateway/api/middleware"
	"api-gateway/config/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connR, connP *grpc.ClientConn, logger logger.Logger) *gin.Engine {
	h := handlers.NewHandler(connR, connP, logger)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())

	// Reservation routes
	reservation := protected.Group("/reservation")
	reservation.POST("/", h.ReservationCreate)
	reservation.GET("/:id", h.ReservationGet)
	reservation.PUT("/:id", h.ReservationUpdate)
	reservation.DELETE("/:id", h.ReservationDelete)
	protected.GET("/reservations", h.ReservationGetAll)
	reservation.POST("/reservations/check", h.ReservationCheck)
	reservation.POST("/reservations/:id/order")
	reservation.POST("/reservations/:id/payment")

	// Reservation-order routes
	reservationOrder := protected.Group("/reservation_order")
	reservationOrder.POST("/", h.ReservationOrderCreate)
	reservationOrder.GET("/:id", h.ReservationOrderGet)
	reservationOrder.PUT("/:id", h.ReservationOrderUpdate)
	reservationOrder.DELETE("/:id", h.ReservationOrderDelete)
	protected.GET("/reservation_orders", h.ReservationOrderGetAll)

	// Restaurant routes
	restaurant := protected.Group("/restaurant")
	restaurant.POST("/", h.RestaurantCreate)
	restaurant.GET("/:id", h.RestaurantGet)
	restaurant.PUT("/:id", h.RestaurantUpdate)
	restaurant.DELETE("/:id", h.RestaurantDelete)
	protected.GET("/restaurants", h.RestaurantGetAll)

	// Menu routes
	menu := protected.Group("/menu")
	menu.POST("/", h.MenuCreate)
	menu.GET("/:id", h.MenuGet)
	menu.PUT("/:id", h.MenuUpdate)
	menu.DELETE("/:id", h.MenuDelete)
	protected.GET("/menus", h.MenuGetAll)

	// Payment routes
	payment := protected.Group("/payment")
	payment.GET("/:id", h.GetPaymentByID)
	payment.PUT("/:id", h.UpdatePayment)

	return router
}
