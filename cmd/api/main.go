package main

import (
	"github.com/M1ralai/me-portfolio/cmd/api/db"
	"github.com/M1ralai/me-portfolio/cmd/api/server"
	"github.com/M1ralai/me-portfolio/cmd/api/types"
)

func main() {
	db.InitDb()
	s := server.NewServer(":3000", ":8080", types.NewLogger("server"))
	s.Run()
}
