package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	database := router.Group("/database")
	{
		database.GET("/all", h.SelectAll)
		orderBy := database.Group("/orderBy")
		{
			orderBy.GET("/id", h.SelectOrderByID)
		}
	}

	return router
}
