package server

import (
	"giftCard/controller"
	"giftCard/middlewares"
	"github.com/gin-gonic/gin"
)

func Run() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middlewares.Cors())

	//router.GET("/", func(c *gin.Context) {
	//	c.Writer.WriteString("hello")
	//})

	router.GET("/api/v1/contacts", controller.AllContact)
	router.POST("/api/v1/contact/create", controller.AddContact)
	router.GET("/api/v1/contact/delete", controller.DeleteContact)

	router.GET("/api/v1/categories", controller.GetCategories)
	router.POST("/api/v1/category/create", controller.CreateCategory)
	router.GET("/api/v1/category/delete", controller.DeleteCategory)
	router.POST("/api/v1/category/update", controller.UpdateCategoryById)

	router.GET("/api/v1/cards", controller.GetCardsByCategory)
	router.POST("/api/v1/card/create", controller.CreateCard)
	router.GET("/api/v1/card/delete", controller.DeleteCard)
	router.POST("/api/v1/card/update", controller.UpdateCardById)

	router.GET("/api/v1/sell", controller.SellCard)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
