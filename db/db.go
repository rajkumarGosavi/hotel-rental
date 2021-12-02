package db

import (
	"context"
	"hotel-rental/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"

type HotelStorer interface {
	// hotel
	CreateHotel(ctx context.Context, hotel Hotel) (err error)
	GetHotels(ctx context.Context) (hotels []Hotel, err error)
	GetHotelDetails(ctx context.Context, hotelID int64) (details HotelDetails, err error)
}

type RoomStorer interface {
	// room
	GetRoom(ctx context.Context, roomID, hotelID int64) (room Room, err error)
	RentRoom(ctx context.Context, req BookRoomRequest) (err error)
	CreateRoom(ctx context.Context, room Room) (err error)
	SlotAvailability(ctx context.Context, roomID, hotelID int64, from, to string) (available bool, err error)
}

type Storer interface {
	HotelStorer
	RoomStorer
}

type pgStore struct {
	db *sqlx.DB
}

func Init() (s Storer, err error) {
	uri := config.ReadEnvString("DB_URI")
	conn, err := sqlx.Connect(dbDriver, uri)
	if err != nil {
		log.Println("err", err.Error())
		return
	}
	s = &pgStore{conn}

	return
}
