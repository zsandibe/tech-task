package sendEmail

import (
	// Импортируйте библиотеку mimetype
	"fmt"
	"net/http"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func (s *SendEmailHandler) SendEmail(c *gin.Context) {
	// Получаем файл и список почт из запроса
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	emailsStr := c.PostForm("emails")
	emails := strings.Split(emailsStr, ",")

	// Ограничение на типы файлов
	allowedMimeTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}

	fileHeader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		return
	}

	mime, err := mimetype.DetectReader(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to detect MIME type"})
		return
	}
	fmt.Println(mime)
	if !allowedMimeTypes[mime.String()] {
		fmt.Println("OK")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	// Отправляем файл по почте
	if err := s.send.SendFileToEmail(emails, fileHeader, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
