package main

import (
	"github.com/cyril_pierro/Rest_api_2/helpers"
	"github.com/cyril_pierro/Rest_api_2/router"
)

func main() {
	helpers.StartDb()
	router.StartRoutes()
}
