package handlers

import (
	"api-gateway/config/logger"
	pbp "api-gateway/genproto/payment"
	pbr "api-gateway/genproto/reservation"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Reservation      pbr.ReservationServiceClient
	Restaurant       pbr.RestaurantServiceClient
	Payment          pbp.PaymentServiceClient
	Menu             pbr.MenuServiceClient
	ReservationOrder pbr.ReservationOrderServiceClient
	Logger           logger.Logger
}

func NewHandler(connR, connP *grpc.ClientConn, l logger.Logger) *HTTPHandler {
	return &HTTPHandler{
		Reservation:      pbr.NewReservationServiceClient(connR),
		Restaurant:       pbr.NewRestaurantServiceClient(connR),
		Payment:          pbp.NewPaymentServiceClient(connP),
		Menu:             pbr.NewMenuServiceClient(connR),
		ReservationOrder: pbr.NewReservationOrderServiceClient(connR),
		Logger:           l,
	}
}
