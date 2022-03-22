package main

import "NetworkDisk/router"

func main() {

	r := router.SetupRouter()
	r.Run(":8081")

}