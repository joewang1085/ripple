package main

import (
	"controllers"

	"github.com/ripple"
)

func main() {

	// Build the REST application
	app := ripple.NewApplication()
	// Create a controller and register it. Any number of controllers
	// can be registered that way.
	appController := controllers.NewAppController()
	app.RegisterController("appName", appController)
	// Setup the routes. The special patterns `_controller` will automatically match
	// an existing controller, as defined above. Likewise, `_action` will match any
	// existing action.
	app.AddRoute(ripple.Route{Pattern: ":_controller/:_action/:params"})
	app.AddRoute(ripple.Route{Pattern: ":_controller/:_action"})

	fmt.Println("web listening 127.0.0.1:8881...")
	http.ListenAndServe(":8881", app)
}
