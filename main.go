package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/routers"
)

func main() {
	router := gin.Default()
	routers.InitRouter(router)
	router.Static("/views", "./views")
	router.LoadHTMLGlob("./views/html/*")
	//start serve
	router.Run(":8000")
}
