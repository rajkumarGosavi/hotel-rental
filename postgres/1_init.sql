CREATE TABLE IF NOT EXISTS hotels (id serial PRIMARY KEY, name varchar(255), address varchar(255));
CREATE TABLE IF NOT EXISTS rooms (id serial PRIMARY KEY, hotel_id int, rate_per_hour varchar(11), CONSTRAINT fk_hotels FOREIGN KEY (hotel_id) REFERENCES hotels(id));
CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, name varchar(255));
CREATE TABLE IF NOT EXISTS bookings (id serial PRIMARY KEY, room_id int, user_id int, rented_from timestamp, rented_to timestamp, CONSTRAINT fk_rooms FOREIGN KEY (room_id) REFERENCES rooms(id), CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id));
INSERT INTO users (name) VALUES ('test_user');