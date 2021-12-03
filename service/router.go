package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	router.HandleFunc("/hotel", CreateHotel(deps)).Methods(http.MethodPost)
	router.HandleFunc("/hotel/{id}", GetHotelDetails(deps)).Methods(http.MethodGet)
	router.HandleFunc("/hotels", ListHotels(deps)).Methods(http.MethodGet)

	router.HandleFunc("/room", CreateRoom(deps)).Methods(http.MethodPost)
	router.HandleFunc("/room/{hotelID}/{id}", GetRoom(deps)).Methods(http.MethodGet)
	router.HandleFunc("/room/{hotelID}/{id}", RentRoom(deps)).Methods(http.MethodPut)
	router.HandleFunc("/bookings/{id}", GetRoomBookings(deps)).Methods(http.MethodGet)

	return
}
