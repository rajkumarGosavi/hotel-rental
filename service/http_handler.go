package service

import (
	"encoding/json"
	"hotel-rental/db"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateHotel(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		var hotelData db.Hotel
		rawData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("failed to read from request", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(rawData, &hotelData)
		if err != nil {
			log.Println("failed to unmarshal request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.HotelService.CreateHotel(req.Context(), hotelData)
		if err != nil {
			log.Println("failed to create hotel", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Success"))
	})
}

func GetHotelDetails(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		vars := mux.Vars(r)
		hotelID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		details, err := deps.HotelService.GetHotelDetails(r.Context(), hotelID)
		if err != nil {
			log.Println("err", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(details)
		if err != nil {
			log.Println("err", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(res)
	})
}

func ListHotels(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		hotels, err := deps.HotelService.GetHotels(r.Context())
		if err != nil {
			log.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rawHotels, err := json.Marshal(hotels)
		if err != nil {
			log.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Write(rawHotels)
		rw.WriteHeader(http.StatusOK)
	})
}
func CreateRoom(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var room db.RoomRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read body", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &room)
		if err != nil {
			log.Println("failed to read body", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var req db.Room
		req.HotelID = room.HotelID
		req.RatePerHour = room.RatePerHour

		err = deps.RoomService.CreateRoom(r.Context(), req)
		if err != nil {
			log.Println("failed to create room", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Success"))
	})
}

func GetRoom(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hotelID, err := strconv.ParseInt(vars["hotelID"], 10, 64)
		if err != nil {
			log.Println("invalid request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid hotel id"))
			return
		}
		roomID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Println("invalid request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid room id"))
			return
		}

		roomData, err := deps.RoomService.GetRoom(r.Context(), roomID, hotelID)
		if err != nil {
			log.Println("failed to get room details", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed to get room data"))
			return
		}

		rawRoomData, err := json.Marshal(roomData)
		if err != nil {
			log.Println("failed to read room details", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed to read room data"))
			return
		}
		rw.Write(rawRoomData)
		rw.WriteHeader(http.StatusOK)
	})
}

func RentRoom(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hotelID, err := strconv.ParseInt(vars["hotelID"], 10, 64)
		if err != nil {
			log.Println("invalid request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid hotel id"))
			return
		}
		roomID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Println("invalid request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid room id"))
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("failed to read request"))
			return
		}

		var rentReq db.BookRoom
		err = json.Unmarshal(body, &rentReq)
		if err != nil {
			log.Println("failed to read request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("failed to read request"))
			return
		}

		from, err := time.Parse(layout, rentReq.RentedFrom)
		if err != nil {
			log.Println("failed to parse rented from", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("failed to parse rented from"))
			return
		}

		to, err := time.Parse(layout, rentReq.RentedTo)
		if err != nil {
			log.Println("failed to parse rented to", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("failed to parse rented to"))
			return
		}

		status, err := deps.RoomService.RentRoom(r.Context(), db.BookRoomRequest{
			RoomID:  roomID,
			HotelID: hotelID,
			From:    &from,
			To:      &to,
			UserID:  rentReq.UserID,
		})
		if err != nil {
			log.Println("failed to rent room details", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed to rent room data"))
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(status))
	})
}

func GetRoomBookings(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Println("invalid request", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid room id"))
			return
		}

		roomBookings, err := deps.RoomService.GetRoomBookings(r.Context(), roomID)
		if err != nil {
			log.Println("failed to get room details", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed to get room data"))
			return
		}

		rawRoomBookings, err := json.Marshal(roomBookings)
		if err != nil {
			log.Println("failed to read room details", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("failed to read room data"))
			return
		}
		rw.Write(rawRoomBookings)
		rw.WriteHeader(http.StatusOK)
	})
}
