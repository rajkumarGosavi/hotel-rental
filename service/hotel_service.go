package service

import (
	"context"
	"hotel-rental/db"
)

type HotelServicer interface {
	CreateHotel(ctx context.Context, req db.Hotel) (err error)
	GetHotelDetails(ctx context.Context, hotelID int64) (rooms db.HotelDetails, err error)
	GetHotels(ctx context.Context) (hotels []db.Hotel, err error)
}

type hotelService struct {
	Store db.Storer
}

func NewHotelService(store db.Storer) HotelServicer {
	return &hotelService{
		Store: store,
	}
}

func (s *hotelService) CreateHotel(ctx context.Context, req db.Hotel) (err error) {
	return s.Store.CreateHotel(ctx, req)
}

func (s *hotelService) GetHotelDetails(ctx context.Context, hotelID int64) (rooms db.HotelDetails, err error) {
	return s.Store.GetHotelDetails(ctx, hotelID)
}
func (s *hotelService) GetHotels(ctx context.Context) (hotels []db.Hotel, err error) {
	return s.Store.GetHotels(ctx)
}
