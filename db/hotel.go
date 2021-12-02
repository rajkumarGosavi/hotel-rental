package db

import (
	"context"
)

type Hotel struct {
	ID      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}

type HotelDetails struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Rooms   []Room `json:"rooms"`
}

func (s *pgStore) CreateHotel(ctx context.Context, hotel Hotel) (err error) {
	_, err = s.db.ExecContext(ctx, createHotel, hotel.Name, hotel.Address)
	return
}

func (s *pgStore) GetHotelDetails(ctx context.Context, hotelID int64) (details HotelDetails, err error) {
	var hotel Hotel
	err = s.db.GetContext(ctx, &hotel, getHotel, hotelID)
	if err != nil {
		return
	}
	rooms, err := s.getHotelRooms(ctx, hotelID)
	details = HotelDetails{
		ID:      hotel.ID,
		Name:    hotel.Name,
		Address: hotel.Address,
		Rooms:   rooms,
	}
	return
}

func (s *pgStore) GetHotels(ctx context.Context) (hotels []Hotel, err error) {
	err = s.db.SelectContext(ctx, &hotels, getHotels)
	return
}

func (s *pgStore) getHotelRooms(ctx context.Context, hotelID int64) (rooms []Room, err error) {
	err = s.db.SelectContext(ctx, &rooms, getHotelRooms, hotelID)
	return
}
