package data

import (
    "encoding/xml"
)

// northbound struct
type Northbound_request struct {
	ActionID         string           `json:"actionid"`
	Target           string           `json:"target"`
	ActionDefinition Actiondefinition `json:"actiondefinition"`
}

type Actiondefinition struct {
	ActionType string `json:"actiontype"`
	Service    string `json:"service"`
	Action     string `json:"action"`
}

type Northbound_response struct {
	ActionID string `json:"actionid"`
	Status   string `json:"status"`
}

// southbound struct
type Southbound_request struct {
	XMLName   xml.Name `xml:"ITResourceOrchestration"`
    ITResource ITResource `xml:"ITResource"`
}

type ITResource struct {
	configuration configurationXML `xml:"configuration"`
}

type configurationXML struct {
	capability string `xml:"capability"`
	configurationRule string `xml:"configurationRule"`
	Name string `xml:"Name"`
}