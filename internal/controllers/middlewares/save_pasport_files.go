package middlewares

import (
	"ecosnap/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SavePassportFiles(c *gin.Context) {
	files := map[string]string{
		"front_side_of_the_passport_file": FrontSideOfThePassportPath,
		"back_side_of_the_passport_file":  BackSideOfThePassportPath,
		"selfie_with_passport_file":       SelfieWithPassportPath,
	}

	for formKey, contextKey := range files {
		file, err := c.FormFile(formKey)
		if err != nil {
			logger.Error.Print(fmt.Sprintf("Missing file: %s", formKey))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Missing file: %s", formKey),
			})
			return
		}

		// Определяем расширение файла (без точки)
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
		if ext == "" {
			logger.Error.Print(fmt.Sprintf("File %s has no extension", file.Filename))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("File %s has no extension", file.Filename),
			})
			return
		}

		// Собираем путь: uploads/<ext>/
		dirPath := filepath.Join("uploads", ext)

		// Создаём папку, если не существует
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			logger.Error.Print(fmt.Sprintf("Could not create directory: %s", dirPath))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not create directory: %s", dirPath),
			})
			return
		}

		// Уникальное имя файла
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
		savePath := filepath.Join(dirPath, fileName)

		// Сохраняем файл
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			logger.Error.Print(fmt.Sprintf("Could not save file: %s", formKey))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not save file: %s", formKey),
			})
			return
		}

		// Кладём путь в контекст
		c.Set(contextKey, savePath)
	}

	c.Next()
}
