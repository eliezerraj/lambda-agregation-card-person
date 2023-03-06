package domain

type AgregationCardPerson struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	CardNumber		string  `json:"card_number,omitempty"`
	Person			string  `json:"person,omitempty"`
	Status			string  `json:"status,omitempty"`
	Valid			string  `json:"valid,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewAgregationCardPerson(id string, 
			sk string, 
			cardnumber string, 
			person string,
			status	string,
			valid	string,
			tenant	string) *AgregationCardPerson{
	return &AgregationCardPerson{
		ID:	id,
		SK:	sk,
		CardNumber: cardnumber,
		Person: person,
		Status: status,
		Valid: valid,
		Tenant: tenant,
	}
}