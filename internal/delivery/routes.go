package delivery

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() gin.IRoutes {
	router := gin.Default()

	router.Group("api/archive")
	{
		router.POST("/create")
	}
}
