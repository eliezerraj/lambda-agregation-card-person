package handler

import(
	"time"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/lambda-agregation-card-person/internal/core/domain"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
	"github.com/lambda-agregation-card-person/internal/erro"
	"github.com/lambda-agregation-card-person/internal/service"

)

var childLogger = log.With().Str("handler", "AgregationService").Logger()

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

func (h *AgregationHandler) AddAgregation(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("AddAgregation")

    var agregation domain.AgregationCardPerson
    if err := json.Unmarshal([]byte(req.Body), &agregation); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	time_now := time.Now()
	agregation.CreateAt = &time_now
	response, err := h.agregationService.AddAgregation(agregation)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *AgregationHandler) SetAgregationStatus(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("SetAgregationStatus")

    var agregation domain.AgregationCardPerson
    if err := json.Unmarshal([]byte(req.Body), &agregation); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.agregationService.AddAgregation(agregation)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *AgregationHandler) GetAgregation(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetAgregation")

	id := req.PathParameters["id"]
	sk := req.PathParameters["sk"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

	time_nil := time.Time{}
 	agregation := domain.NewAgregationCardPerson(id,
												sk,
												"",
												"",
												"",
												&time_nil,
												&time_nil,
												"TENANT-001")

	response, err := h.agregationService.GetAgregation(*agregation)
	if err != nil {
		return ApiHandlerResponse(http.StatusNotFound, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}