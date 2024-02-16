package main

import (
	"latihan-gin/src/config"
	"latihan-gin/src/helpers"
	"latihan-gin/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helpers.Migration()
	defer config.DB.Close()
	routes.Router()
}
