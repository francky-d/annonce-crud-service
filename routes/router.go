package routes

import (
	"github.com/franck-djacoto/announce-service/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit(){
	router := gin.Default()

	router.GET("/", controllers.Hello)
	router.Run(":8000")
}

