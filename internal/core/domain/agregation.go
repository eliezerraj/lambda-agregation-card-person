package domain

import (
	"time"

)

type AgregationCardPerson struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	CardNumber		string  `json:"card_number,omitempty"`
	Person			string  `json:"person,omitempty"`
	Status			string  `json:"status,omitempty"`
	CreateAt	 time.Time 	`json:"create_at,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewAgregationCardPerson(id string, 
			sk 			string, 
			cardnumber 	string, 
			person 		string,
			status		string,
			createAt	time.Time,
			tenant	string) *AgregationCardPerson{
	return &AgregationCardPerson{
		ID:	id,
		SK:	sk,
		CardNumber: cardnumber,
		Person: person,
		Status: status,
		CreateAt: createAt, 
		Tenant: tenant,
	}
}