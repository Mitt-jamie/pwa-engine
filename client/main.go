package main

import (
	"pwa-engine/middleware"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	// Register all routes for client-side navigation
	middleware.RegisterRoutes()

	// Start the app in the browser
	app.RunWhenOnBrowser()
}
