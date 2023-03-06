package repository

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/lambda-agregation-card-person/internal/erro"

)

var childLogger = log.With().Str("repository", "AgregationRepository").Logger()

type AgregationRepository struct {
	client 		dynamodbiface.DynamoDBAPI
	tableName   *string
}

func NewAgregationRepository(tableName string) (*AgregationRepository, error){
	childLogger.Debug().Msg("NewAgregationRepository")
	
	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrCreateSession
	}

	return &AgregationRepository{
		client: dynamodb.New(awsSession),
		tableName: aws.String(tableName),
	},nil
}