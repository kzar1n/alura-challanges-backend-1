package main

import (
	"github.com/kzar1n/alura-challanges-backend-1/internal/infra/database"
	"github.com/kzar1n/alura-challanges-backend-1/internal/routes"
)

func main() {
	database.OpenDB()
	routes.HandleRequests()
}
