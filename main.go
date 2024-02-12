package main

import "advertisement_api/api/routes"

func main() {
	r := routes.InitRoutes()
	err := r.Run()

	if err != nil {
		return
	}
}
