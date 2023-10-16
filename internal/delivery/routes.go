package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		if c.Request.Method != "POST" {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": http.StatusMethodNotAllowed})
			c.Abort()
			return
		}
		c.Next()
	})

	api := router.Group("/api")
	{
		archive := api.Group("/archive")
		{
			archive.POST("/info", h.Info.GetInfoByArchive)
			archive.POST("/create", h.Create.CreateArchiveByFiles)

		}
	}

	return router
}
