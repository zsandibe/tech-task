package sendEmail

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *SendEmailHandler) SendEmail(c *gin.Context) {
	// Get the file from the request
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get the file"})
		return
	}
	defer file.Close()

	// Check the MIME type of the file
	mime := fileHeader.Header.Get("Content-Type")
	if mime != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" &&
		mime != "application/pdf" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	// Get the list of emails from the request
	emails := c.PostFormArray("emails")

	// Create a buffer to store the file data
	fileData := bytes.Buffer{}
	if _, err := io.Copy(&fileData, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file data"})
		return
	}

	// Send the email
	err = h.send.SendFileToEmail(fileData.Bytes(), emails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
