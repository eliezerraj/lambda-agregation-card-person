package repository

import(

	"github.com/lambda-agregation-card-person/internal/core/domain"
	"github.com/lambda-agregation-card-person/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func (r *AgregationRepository) Ping() (bool, error){
	return true, nil
}

func (r *AgregationRepository) AddAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("AddAgregation")

	item, err := dynamodbattribute.MarshalMap(agregation)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: r.tableName,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrInsert
	}

	_, err = r.client.TransactWriteItems(transaction)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrInsert
	}

	return &agregation ,nil
}

func (r *AgregationRepository) GetAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("GetAgregation")

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(agregation.ID)),
		expression.Key("sk").BeginsWith(agregation.SK),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.QueryInput{
			TableName:                 r.tableName,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.Query(key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrQuery
	}

	agregation_result := []domain.AgregationCardPerson{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &agregation_result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(agregation_result) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &agregation_result[0], nil
	}
}