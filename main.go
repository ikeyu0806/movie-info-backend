package main

import (
	"github.com/ikeyu0806/movie-info-backend/db"
	"github.com/ikeyu0806/movie-info-backend/server"
)

func main() {
	db.Init()
	server.Init()
}
