package main

import "gin_gonic_products_rest_api/routes"

func main() {
	var r routes.Routes
	r.CreateConnection()
	r.Routes()
	r.Run()
}
