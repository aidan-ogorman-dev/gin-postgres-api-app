#!/bin/bash

export DB_NAME=diary_app
export DB_USER="postgres"
export DB_PASSWORD="postgres"
export DB_HOST=localhost
export DB_PORT="5432"

go build .
go run ./main.go
