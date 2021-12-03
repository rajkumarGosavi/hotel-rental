package service

import (
	"context"
	"errors"
	"hotel-rental/db"
	"log"
)

type RoomServicer interface {
	CreateRoom(ctx context.Context, room db.Room) (err error)
	GetRoom(ctx context.Context, roomID, hotelID int64) (room db.Room, err error)
	RentRoom(ctx context.Context, req db.BookRoomRequest) (status db.RoomStatus, err error)
	GetRoomBookings(ctx context.Context, roomID int64) (bookings []db.Booking, err error)
}

type roomService struct {
	Store db.RoomStorer
}

func NewRoomService(store db.RoomStorer) RoomServicer {
	return &roomService{
		Store: store,
	}
}

func (s *roomService) CreateRoom(ctx context.Context, room db.Room) (err error) {
	err = s.Store.CreateRoom(ctx, room)
	return
}

func (s *roomService) GetRoom(ctx context.Context, roomID, hotelID int64) (room db.Room, err error) {
	room, err = s.Store.GetRoom(ctx, roomID, hotelID)
	return
}

func (s *roomService) RentRoom(ctx context.Context, req db.BookRoomRequest) (status db.RoomStatus, err error) {
	room, err := s.GetRoom(ctx, req.RoomID, req.HotelID)
	if err != nil {
		log.Println("room does not exists")
		return
	}

	isAvailable, err := s.Store.SlotAvailability(ctx, req.RoomID, req.From.Format(layout), req.To.Format(layout))
	if err != nil {
		log.Println("failed to check slot availability", err.Error())
		return
	}

	if !isAvailable {
		log.Println("slot not available")
		err = errors.New(db.ErrSlotNotAvailable)
		return
	}

	err = s.Store.RentRoom(ctx, db.BookRoomRequest{
		RoomID: room.ID,
		From:   req.From,
		To:     req.To,
		UserID: req.UserID,
	})
	status = db.RentedSuccessfully
	return
}

func (s *roomService) GetRoomBookings(ctx context.Context, roomID int64) (bookings []db.Booking, err error) {
	bookings, err = s.Store.GetRoomBookings(ctx, roomID)
	if err != nil {
		log.Println("failed to get room bookings", err.Error())
		return
	}
	return
}
