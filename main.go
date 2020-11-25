package main

import "github.com/eren-dhoheiri/go-rest/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":9900"))
}
