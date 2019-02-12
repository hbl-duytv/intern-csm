package main

import (
	"github.com/hbl-duytv/intern-csm/routers"
)

func main() {
	router := routers.InitRouter()
	router.Static("/views", "./views")
	router.LoadHTMLGlob("./views/html/*")
	//start serve
	router.Run(":8000")
}
