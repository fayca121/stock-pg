package routes

import (
	"database/sql"
	"github.com/fayca121/stock-pg/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *gin.Engine {
	route := gin.Default()

	// Provide db to controllers
	route.Use(func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	})

	apiPathGrp := route.Group("/api/stock")
	{
		apiPathGrp.GET("/:id", controllers.GetStock)
		apiPathGrp.GET("", controllers.GetAllStock)
		apiPathGrp.POST("", controllers.CreateStock)
		apiPathGrp.PUT("/:id", controllers.UpdateStock)
		apiPathGrp.DELETE("/:id", controllers.DeleteStock)
	}

	return route
}
