package main

import "cc-backend-root-console/app/boot"

func main() {
	// Create an instance of the application
	application := boot.BuildApplicationInstance()

	// Boot the application
	application.Boot()
}
