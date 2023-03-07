package service

import(
	"testing"
	"github.com/rs/zerolog"
	"github.com/google/go-cmp/cmp"

	"github.com/lambda-agregation-card-person/internal/core/domain"
	"github.com/lambda-agregation-card-person/internal/repository"
)

var (
	tableName = "agregation_card_person"
	agregationRepository	*repository.AgregationRepository

	agre01 = domain.NewAgregationCardPerson("AGREGATION-01",
											"AGREGATION-01",
											"4444.000.000.001",
											"PERSON-01",
											"ACTIVE",
											"02/26",
											"TENANT-001")
							
)

func TestAddAgregation(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddAgregation Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	result, err := service.AddAgregation(*agre01)
	if err != nil {
		t.Errorf("Error -TestAddAgregation Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(agre01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddAgregation input : %v" , *agre01)
	}
}

func TestGetAgregation(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetAgregation Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	result, err := service.GetAgregation(*agre01)
	if err != nil {
		t.Errorf("Error -TestGetAgregation Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(agre01, result)) {
		t.Logf("Success on TestGetAgregation!!! result : %v ", result)
	} else {
		t.Errorf("Error TestGetAgregation input : %v" , *agre01)
	}
}

func TestSetAgregationStatus(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestSetAgregationStatus Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)

	agre01.Status = "CANCELED"
	result, err := service.SetAgregationStatus(*agre01)
	if err != nil {
		t.Errorf("Error -TestSetAgregationStatus Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(agre01, result)) {
		t.Logf("Success on TestSetAgregationStatus!!! result : %v ", result)
	} else {
		t.Errorf("Error TestSetAgregationStatus input : %v" , *agre01)
	}
}