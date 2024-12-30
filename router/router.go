package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
