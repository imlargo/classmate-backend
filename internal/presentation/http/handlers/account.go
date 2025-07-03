package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	requestsdto "github.com/imlargo/butter/internal/application/dto/requests"
	"github.com/imlargo/butter/internal/presentation/http/responses"
)

type CompletionController interface {
	CreateCompletion(c *gin.Context)
}

type completionControllerImpl struct {
}

func NewCompletionController() CompletionController {
	return &completionControllerImpl{}
}

type CreateCompletionRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content []struct {
			FollowerCnt  int    `json:"follower_count"`
			AccountUrl   string `json:"external_url"`
			HDProfilePic struct {
				URL string `json:"url"`
			} `json:"hd_profile_pic_url_info"`
		} `json:"content"`
	} `json:"messages"`
}

// @Summary 		Create Completion
// @Router			/process [post]
// @Description	Create an Completion
// @Tags			completions
// @Produce		json
// @Param payload body requestsdto.CreateCompletionRequest true "Problem data"
// @Success		200	{object}	string	"Completion"
// @Failure		400	{object}	responses.ErrorResponse	"Bad Request"
// @Failure		500	{object}	responses.ErrorResponse	"Internal Server Error
func (h *completionControllerImpl) CreateCompletion(c *gin.Context) {
	var payload requestsdto.CreateCompletionRequest

	err := c.BindJSON(&payload)
	if err != nil {
		responses.ErrorBadRequest(c, "Parametros invalidos")
	}

	url := "https://api.openai.com/v1/chat/completions"

	prompt := "Solve the following problem. Answer only with the correct solution and a very short explanation, written in Spanish. Make sure the answer is accurate, and keep the explanation as brief and clear as possible. Give the option number and the correct solution on the first line with bold, Then, on a second line, give the explanation"
	requestBody := map[string]interface{}{
		"prompt": prompt,
	}

	var body io.Reader
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		responses.ErrorBadRequest(c, "error marshaling JSON: "+err.Error())
		return
	}

	body = bytes.NewBuffer(jsonData)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, body)
	if err != nil {
		responses.ErrorBadRequest(c, "error creating request: "+err.Error())
		return
	}

	apiKey := "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"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", apiKey),
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		responses.ErrorInternalServerWithMessage(c, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		responses.ErrorInternalServerWithMessage(c, string(bodyBytes))
		return
	}

	var responseBody any
	if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
		responses.ErrorInternalServerWithMessage(c, err.Error())
		return
	}

	c.JSON(200, "Okaaa")
}
