package data

// northbound struct
type northbound_request struct {
	ActionID         string           `json:"actionid"`
	Target           string           `json:"target"`
	ActionDefinition actiondefinition `json:"actiondefinition"`
}

type actiondefinition struct {
	ActionType string `json:"actiontype"`
	Service    string `json:"service"`
	Action     string `json:"action"`
}

type northbound_response struct {
	ActionID string `json:"actionid"`
	Status   string `json:"status"`
}
