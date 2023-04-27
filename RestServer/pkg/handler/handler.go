package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		database := api.Group("/database")
		{
			database.GET("/all", h.SelectAll)
			orderBy := database.Group("/orderBy")
			{
				orderBy.GET("/id", h.SelectOrderByID)
			}
		}

		user := api.Group("/user")
		{
			user.POST("/logon", h.Logon)
			user.POST("/login", h.Login)
		}
	}

	return router
}
