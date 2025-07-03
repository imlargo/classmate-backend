package handlers

import (
	"github.com/gin-gonic/gin"
	requestsdto "github.com/imlargo/butter/internal/application/dto/requests"
	"github.com/imlargo/butter/internal/application/services"
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

// @Summary 		Create Completion from image
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

	completion, err := services.GetCompletionForImage(payload.Image)
	if err != nil {
		responses.ErrorInternalServerWithMessage(c, "Error al generar la respuesta: "+err.Error())
		return
	}

	c.JSON(200, completion)
}
