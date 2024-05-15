package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"github.com/rsj-rishabh/urbanClapClone/server/app"
	"github.com/rsj-rishabh/urbanClapClone/server/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.DBMigrate()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"}, // Allow requests from localhost:4200
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          false,
	})

	handler := c.Handler(app.Router)

	fmt.Println("Working server on port:3000")
	http.ListenAndServe(":3000", handler)
}
