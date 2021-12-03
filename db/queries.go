package db

const (
	createHotel   = `INSERT INTO hotels (name, address) VALUES ($1, $2)`
	getHotel      = `SELECT * FROM hotels WHERE id = $1`
	getHotels     = `SELECT * FROM hotels`
	getHotelRooms = `SELECT * FROM rooms WHERE hotel_id=$1`
	createRoom    = `INSERT INTO rooms (hotel_id, rate_per_hour)
		VALUES ($1, $2)`
	getRoom  = `SELECT * FROM rooms WHERE id=$1 AND hotel_id=$2`
	rentRoom = `INSERT INTO bookings  (room_id, user_id, rented_from, rented_to)
		VALUES ($1, $2, $3, $4)`
	// updateIsRentedForRoom = `UPDATE rooms SET is_rented=$3 WHERE id=$1 AND hotel_id=$2`
	slotAvailability = `SELECT * FROM bookings WHERE room_id=$1 AND (rented_to >= $2 AND rented_from <= $3);`
	getRoomBookings  = `SELECT * FROM bookings WHERE room_id=$1`
)
