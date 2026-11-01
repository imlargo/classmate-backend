package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiKey      = "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"
	prompt      = "Resuelve el siguiente problema. Si la pregunta tiene una sola respuesta correcta, responde únicamente con el número de la opción (a,b,c,d) o sino con el texto de la opcion de respuesta o respuesta. Si la pregunta tiene varias respuestas correctas, responde únicamente las opciones separadas por coma en una sola linea. No incluyas justificación ni explicaciones adicionales. Solo las respuestas correctas de la forma más compacta posible. Realiza todos los cálculos y verificaciones necesarias con precisión (unidades y fórmulas correctas). Investiga/contrasta conceptos cuando aplique; evita alucinaciones y no inventes datos ni supuestos no dados. Si la información es insuficiente para una respuesta cierta, responde únicamente Información insuficiente"
	openAIModel = "gpt-5-2025-08-07"
)

type ImageURL struct {
	URL string `json:"url"`
}

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type MessageText struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type RequestBodyText struct {
	Model    string        `json:"model"`
	Messages []MessageText `json:"messages"`
}

type Choice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func GetCompletionForImage(imageBase64 string) (string, error) {
	requestBody := RequestBody{
		Model: openAIModel,
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: prompt,
					},
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: imageBase64,
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(result.Choices) > 0 && result.Choices[0].Message.Content != "" {
		return result.Choices[0].Message.Content, nil
	}

	return "No se pudo obtener una respuesta.", nil
}

func GetCompletionForText(imageBase64 string) (string, error) {
	requestBody := RequestBodyText{
		Model: openAIModel,
		Messages: []MessageText{
			{
				Role:    "user",
				Content: fmt.Sprintf("%s\n\n%s", prompt, imageBase64),
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if len(result.Choices) > 0 && result.Choices[0].Message.Content != "" {
		return result.Choices[0].Message.Content, nil
	}

	return "No se pudo obtener una respuesta.", nil
}
