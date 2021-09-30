package routes

import (
	. "github.com/franck-djacoto/announce-service/conf"
	. "github.com/franck-djacoto/announce-service/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit(app *Application) {
	router := gin.Default()
	AnControl := (&AnnonceController{}).New(app.Db)

	router.GET("/", AnControl.ChecIfServiceRespond)
	groupAnn := router.Group("/api/annonce")
	{
		groupAnn.POST("/add", AnControl.Add)
		groupAnn.POST("/update/:id", AnControl.Update)
		groupAnn.GET("/delete/:id", AnControl.Delete)
		groupAnn.GET("/all", AnControl.All)
		groupAnn.GET("/detail/:id", AnControl.Detail)
		groupAnn.POST("/search", AnControl.Search)
	}

	router.Run(":8000")
}
