package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type RoomStatus string

const (
	IsRented           RoomStatus = "Room already rented"
	RentedSuccessfully RoomStatus = "Room has been rented successfully"
)
const (
	ErrAlreadyRented    = "room already rented"
	ErrSlotNotAvailable = "slot not available"
)

type Room struct {
	ID          int64  `db:"id"`
	HotelID     int64  `json:"hotel_id" db:"hotel_id"`
	RatePerHour string `json:"rate_per_hour" db:"rate_per_hour"`
}

type RoomRequest struct {
	HotelID     int64  `json:"hotel_id"`
	RatePerHour string `json:"rate_per_hour"`
}

type BookRoom struct {
	UserID     int64  `json:"user_id"`
	RentedFrom string `json:"rented_from"`
	RentedTo   string `json:"rented_to"`
}

type BookRoomRequest struct {
	RoomID  int64
	HotelID int64
	UserID  int64
	From    *time.Time
	To      *time.Time
}

type Booking struct {
	ID      int64  `json:"id,omitempty" db:"id"`
	RoomID  int64  `json:"room_id" db:"room_id"`
	HotelID int64  `json:"hotel_id" db:"hotel_id"`
	UserID  int64  `json:"user_id" db:"user_id"`
	From    string `json:"from" db:"rented_from"`
	To      string `json:"to" db:"rented_to"`
}

func (s *pgStore) GetRooms(ctx context.Context, hotelID int64) (rooms HotelDetails, err error) {
	return
}

func (s *pgStore) CreateRoom(ctx context.Context, room Room) (err error) {
	_, err = s.db.ExecContext(ctx, createRoom, room.HotelID, room.RatePerHour)
	return
}
func (s *pgStore) GetRoom(ctx context.Context, roomID int64, hotelID int64) (room Room, err error) {
	err = s.db.GetContext(ctx, &room, getRoom, roomID, hotelID)
	return
}
func (s *pgStore) RentRoom(ctx context.Context, req BookRoomRequest) (err error) {
	_, err = s.db.ExecContext(ctx, rentRoom, req.RoomID, req.UserID, req.From, req.To)
	return
}

// SlotAvailability will check if the slot is available for renting
// the slot is available if
// there are no bookings whose rented_to is less than the new booking start_time
// and rented_form is more than the end_time
func (s *pgStore) SlotAvailability(ctx context.Context, roomID int64, from, to string) (available bool, err error) {
	var slot Booking
	fmt.Println("from", from, "to", to)
	err = s.db.GetContext(ctx, &slot, slotAvailability, roomID, from, to)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			fmt.Println("no rows")
			available = true
			err = nil
			return
		}
		fmt.Println("some err", err.Error())
	}
	return
}
