package archiveCreate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ArchiveCreateHandler) CreateArchiveByFiles(c *gin.Context) {
	// Получить список файлов из запроса
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB - максимальный размер для файлов
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := c.Request.MultipartForm.File["files[]"]

	// Определить допустимые MIME-типы
	validMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
	}

	// Проверить и создать архивы для файлов
	archiveFiles, err := h.create.CreateArchive(files, validMimeTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Отправить архивы как ответ
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=archives.zip")
	c.Data(http.StatusOK, "application/zip", archiveFiles)
}
