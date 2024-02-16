package routes

import (
	"latihan-gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		ping := v1.Group("/ping")
		{
			ping.GET("/", controllers.PingController)
		}
		month := v1.Group("/month")
		{
			month.GET("/list", controllers.MonthList)
			month.GET("/:id", controllers.MonthShow)
			month.POST("/create", controllers.MonthCreate)
			month.PUT("/update/:id", controllers.MonthUpdate)
			month.DELETE("/delete/:id", controllers.MonthDestroy)
		}
	}
	router.Run(":8080")
}
