package main

import "challenge-2-chapter-2/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
