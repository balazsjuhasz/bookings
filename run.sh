#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -dbhost=localhost -dbport=19999 -dbname=bookings -dbuser=bookings -dbpass=bookings -cache=false -production=false