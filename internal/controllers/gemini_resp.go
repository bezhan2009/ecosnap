package controllers

import (
	"ecosnap/internal/app/service/ai"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"mime/multipart"
	"net/http"
)

// Request структура для входящего запроса
type GeminiRequest struct {
	Text  string                `form:"text"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// Response структура для ответа
type GeminiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func AnalyzeImageWithText(c *gin.Context) {
	var request GeminiRequest

	// Парсим multipart/form-data запрос с явным указанием типа
	if err := c.ShouldBindWith(&request, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, GeminiResponse{
			Success: false,
			Error:   "Неверный формат запроса: " + err.Error(),
		})
		return
	}

	request.Text = `Посмотри что на изоброжения потом напиши какой это вид мусора из этих трех категорий: 
recycle

compostable

landfill


в ответе верни только категорию одно слово если не одно из категорий не подходит верни просто цифру 0
`

	// Проверяем, что файл был загружен
	if request.Image == nil {
		c.JSON(http.StatusBadRequest, GeminiResponse{
			Success: false,
			Error:   "Изображение обязательно для загрузки",
		})
		return
	}

	// Открываем загруженный файл
	file, err := request.Image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeminiResponse{
			Success: false,
			Error:   "Ошибка при открытии файла: " + err.Error(),
		})
		return
	}
	defer file.Close()

	// Читаем данные изображения
	imageData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeminiResponse{
			Success: false,
			Error:   "Ошибка при чтении файла: " + err.Error(),
		})
		return
	}

	// Определяем MIME type изображения
	mimeType := http.DetectContentType(imageData)

	// Проверяем поддерживаемые форматы изображений
	supportedMimeTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
		"image/gif":  true,
	}

	if !supportedMimeTypes[mimeType] {
		c.JSON(http.StatusBadRequest, GeminiResponse{
			Success: false,
			Error:   fmt.Sprintf("Неподдерживаемый формат изображения: %s. Поддерживаются: JPEG, PNG, WebP, GIF", mimeType),
		})
		return
	}

	// Отправляем данные в Gemini
	response, err := ai.SendToGemini(request.Text, imageData, mimeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GeminiResponse{
			Success: false,
			Error:   "Ошибка при обращении к Gemini: " + err.Error(),
		})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, GeminiResponse{
		Success: true,
		Message: "Успешно обработано",
		Data:    response,
	})
}
