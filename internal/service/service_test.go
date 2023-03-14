package service

import(
	"testing"
	"time"
	"github.com/rs/zerolog"
	"github.com/google/go-cmp/cmp"

	"github.com/lambda-agregation-card-person/internal/core/domain"
	"github.com/lambda-agregation-card-person/internal/repository"
)

var (
	tableName = "agregation_card_person"
	agregationRepository	*repository.AgregationRepository
							
)

func TestAddAgregation(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewAgregationRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddAgregation Create Repository DynanoDB")
	}

	service	:= NewAgregationService(*repository)
	var time_now = time.Now()

	agre01 := domain.NewAgregationCardPerson("AGREGATION-555.000.000.001",
											"PERSON:PERSON-555",
											"555.000.000.001",
											"MR TEST 555",
											"CANCELED-TEST",
											&time_now,
											&time_now,
											"TENANT-555")

	result, err := service.AddAgregation(*agre01)
	if err != nil {
		t.Errorf("Error -TestAddAgregation Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(agre01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddAgregation input : %v || result : %v " , *agre01, result)
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

	agre01 := domain.NewAgregationCardPerson("AGREGATION-4444.000.000.001",
											"PERSON:PERSON-001",
											"",
											"",
											"",
											nil,
											nil,
											"")

	result, err := service.GetAgregation(*agre01)
	if err != nil {
		t.Errorf("Error -TestGetAgregation Access DynanoDB %v ", tableName)
	}

	if (agre01.SK == result.SK){
		t.Logf("Success on TestGetAgregation!!! result : %v ", result)
	}else {
		t.Errorf("Error TestGetAgregation input : %v || result : %v "  , *agre01, result)
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
	agre01 := domain.NewAgregationCardPerson("AGREGATION-4444.000.000.001",
											"PERSON:PERSON-001",
											"",
											"",
											"CANCELED",
											nil,
											nil,
											"")

	result, err := service.SetAgregationStatus(*agre01)
	if err != nil {
		t.Errorf("Error -TestSetAgregationStatus Access DynanoDB %v ", tableName)
	}

	if (agre01.Status == result.Status){
		t.Logf("Success on TestSetAgregationStatus!!! result : %v ", result)
	}else {
		t.Errorf("Error TestSetAgregationStatus input : %v || result : %v "  , *agre01, result)
	}
}