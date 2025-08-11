package main

import (
	"github.com/upinmcSE/goshop/internal/app"
	"github.com/upinmcSE/goshop/internal/config"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// Start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}