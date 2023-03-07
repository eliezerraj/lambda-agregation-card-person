package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lambda-agregation-card-person/internal/repository"
	"github.com/lambda-agregation-card-person/internal/service"
	"github.com/lambda-agregation-card-person/internal/handler"

)

var (
	logLevel		=	zerolog.DebugLevel // InfoLevel DebugLevel
	tableName		=	"agregation_card_person"
	version			=	"lambda-aggregation_person_card (github) version 1.0"
	eventSource		=	"lambda-card"
	eventBusName	=	"event-bus-card"	
	response			*events.APIGatewayProxyResponse
	agregationHandler		*handler.AgregationHandler
	agregationRepository	*repository.AgregationRepository
	agregationService		*service.AgregationService
)

func getEnv(){
	if os.Getenv("TABLE_NAME") !=  "" {
		tableName = os.Getenv("TABLE_NAME")
	}
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.DebugLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init(){
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main(){
	log.Debug().Msg("main lambda-aggregation_person_card (go) v 2.0")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	agregationRepository, err := repository.NewAgregationRepository(tableName)
	if err != nil{
		return
	}

	agregationService 	= service.NewAgregationService(*agregationRepository)
	agregationHandler 	= handler.NewAgregationHandler(*agregationService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("handler")
	log.Debug().Msg("-------------------")
	log.Debug().Str("req.Body", req.Body).
				Msg("APIGateway Request.Body")
	log.Debug().Msg("--------------------")

	switch req.HTTPMethod {
		case "GET":
			if (req.Resource == "/agregation/{id}/{sk}"){
				response, _ = agregationHandler.GetAgregation(req)
			}else if (req.Resource == "/version"){
				response, _ = agregationHandler.GetVersion(version)
			}else {
				response, _ = agregationHandler.UnhandledMethod()
			}
		case "POST":
			if (req.Resource == "/agregation"){
				response, _ = agregationHandler.AddAgregation(req)
			}else if (req.Resource == "/agregation/status") {
				response, _ = agregationHandler.SetAgregationStatus(req)
			}else {
				response, _ = agregationHandler.UnhandledMethod()
			}
		case "DELETE":
			response, _ = agregationHandler.UnhandledMethod()
		case "PUT":
			response, _ = agregationHandler.UnhandledMethod()
		default:
			response, _ = agregationHandler.UnhandledMethod()
	}

	return response, nil
}