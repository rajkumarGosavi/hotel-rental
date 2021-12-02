package main

import (
	"fmt"
	"hotel-rental/config"
	"hotel-rental/db"
	"hotel-rental/service"
	"log"

	"github.com/urfave/negroni"
)

func init() {
	config.Load()
}

func main() {
	store, err := db.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	hotelService := service.NewHotelService(store)
	roomService := service.NewRoomService(store)
	deps := service.NewDependecies(hotelService, roomService)

	router := service.InitRouter(deps)

	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort()
	addr := fmt.Sprintf(":%d", port)

	server.Run(addr)
}
