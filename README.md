
CREATE DATABASE hotel_rental;

CREATE TABLE IF NOT EXISTS hotels (id serial PRIMARY KEY, name varchar(255), address varchar(255));

CREATE TABLE IF NOT EXISTS rooms (id serial PRIMARY KEY, hotel_id int, rate_per_hour varchar(11),
CONSTRAINT fk_hotels FOREIGN KEY (hotel_id) REFERENCES hotels(id));

CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, name varchar(255));

CREATE TABLE IF NOT EXISTS bookings (id serial PRIMARY KEY, room_id int, user_id int, rented_from timestamp, rented_to timestamp, CONSTRAINT fk_rooms FOREIGN KEY (room_id) REFERENCES rooms(id), CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id));

INSERT INTO users (name) VALUES ('test_user');

<!-- create hotel -->
curl POST -d '{
    "name": "The Continental",
    "address": "JW street"
}' -v -i 'localhost:9090/hotel'

<!-- list hotels -->
curl -X GET -v -i 'localhost:9090/hotels'

<!-- get hotel details -->
curl -X GET -v -i 'localhost:9090/hotel/2'

<!-- create room -->
curl -X POST -d '{
    "hotel_id": 2,
    "rate_per_hour": "60"
}' -v -i 'localhost:9090/room'

<!-- Get room -->
curl -X GET -v -i 'localhost:9090/room/2/1'


<!-- Rent room -->
curl -X PUT -d '{
    "user_id":1,
    "rented_from": "2021-12-03 21:45:00",
    "rented_to": "2021-12-03 21:50:00"
}' -v -i 'localhost:9090/room/2/3'

<!-- Room Bookings -->
curl -X GET -v -i 'localhost:9090/bookings/3'


docker build -t hotel-rental .
docker build -t postgres .
<!-- docker run --rm -p 9090:9090 hotel-rental -->
docker-compose up
docker-compose down