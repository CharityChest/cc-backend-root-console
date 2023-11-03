package main

import (
	"cc-backend-root-console/app/boot/application/application/abstracts"
	"cc-backend-root-console/app/boot/application/application/concrete"
)

func main() {
	// Create an instance of the application
	app := concrete.App().(abstracts.Application)

	// Boot the application
	app.Boot()

	app.Listen()
}
