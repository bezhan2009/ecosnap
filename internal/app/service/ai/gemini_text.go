package ai

import (
	"bytes"
	"ecosnap/internal/app/models"
	"ecosnap/pkg/errs"
	"ecosnap/pkg/logger"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SendToGemini(text string, imageData []byte, mimeType string) (response string, err error) {
	// Создаем базовую структуру запроса
	geminiReq := models.GeminiCandidateReq{
		Contents: []models.GeminiContents{
			{
				Parts: []models.GeminiParts{},
			},
		},
	}

	if text != "" {
		geminiReq.Contents[0].Parts = append(geminiReq.Contents[0].Parts, models.GeminiParts{
			Text: text,
		})
	}

	if len(imageData) > 0 {
		encodedImage := base64.StdEncoding.EncodeToString(imageData)

		geminiReq.Contents[0].Parts = append(geminiReq.Contents[0].Parts, models.GeminiParts{
			InlineData: &models.GeminiInlineData{
				MimeType: mimeType,
				Data:     encodedImage,
			},
		})
	}

	// Проверяем, что есть хотя бы одна часть (текст или изображение)
	if len(geminiReq.Contents[0].Parts) == 0 {
		return "", fmt.Errorf("no content provided: either text or image is required")
	}

	// сериализуем в JSON
	jsonBody, err := json.Marshal(geminiReq)
	if err != nil {
		logger.Error.Printf("[aiService.SendToGemini] Error marshalling json body: %v", err)
		return "", err
	}

	var geminiURL = fmt.Sprintf(
		"%s?key=%s",
		os.Getenv("GEMINI_AI_API"),
		os.Getenv("GEMINI_API_KEY"),
	)

	analyse, err := http.Post(
		geminiURL,
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		logger.Error.Printf("[aiService.SendToGemini] Error sending request: %v", err)
		return "", err
	}
	defer analyse.Body.Close()

	body, err := ioutil.ReadAll(analyse.Body)
	if err != nil {
		logger.Error.Printf("[aiService.SendToGemini] Error reading body: %v", err)
		return "", err
	}

	var GeminiResp models.Gemini
	if err := json.Unmarshal(body, &GeminiResp); err != nil {
		logger.Error.Printf("[aiService.SendToGemini] Error parsing body: %v", err)
		return "", err
	}

	if len(GeminiResp.Candidates) == 0 {
		logger.Error.Printf("[aiService.SendToGemini] No candidates returned from Gemini response: %s", string(body))
		return "", errs.ErrGeminiIsNotWorking
	}
	if len(GeminiResp.Candidates[0].Content.Parts) == 0 {
		logger.Error.Printf("[aiService.SendToGemini] No parts in first candidate's content: %s", string(body))
		return "", errs.ErrGeminiIsNotWorking
	}

	var GeminiText = GeminiResp.Candidates[0].Content.Parts[0].Text

	return GeminiText, nil
}

func SendTextToGemini(text string) (response string, err error) {
	return SendToGemini(text, nil, "")
}

func SendImageToGemini(imageData []byte, mimeType string) (response string, err error) {
	return SendToGemini("", imageData, mimeType)
}
