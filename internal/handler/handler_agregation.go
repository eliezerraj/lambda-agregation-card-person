package handler

import(

	//"encoding/json"

	"github.com/rs/zerolog/log"
	//"github.com/lambda-card/internal/core/domain"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
	"github.com/lambda-agregation-card-person/internal/erro"
	"github.com/lambda-agregation-card-person/internal/service"

)

var childLogger = log.With().Str("handler", "CardHandler").Logger()

var transactionSuccess	= "Transação com sucesso"

type AgregationHandler struct {
	agregationService service.AgregationService
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type MessageBody struct {
	Msg *string `json:"message,omitempty"`
}

func NewAgregationHandler(agregationService service.AgregationService) *AgregationHandler{
	childLogger.Debug().Msg("NewAgregationHandler")
	return &AgregationHandler{
		agregationService: agregationService,
	}
}

func (h *AgregationHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func (h *AgregationHandler) GetVersion(version string) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetVersion")

	response := MessageBody { Msg: &version }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}

	return handlerResponse, nil
}