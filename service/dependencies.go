package service

type Dependencies struct {
	HotelService HotelServicer
	RoomService  RoomServicer
}

func NewDependecies(hotelService HotelServicer, roomService RoomServicer) Dependencies {
	return Dependencies{
		HotelService: hotelService,
		RoomService:  roomService,
	}
}
