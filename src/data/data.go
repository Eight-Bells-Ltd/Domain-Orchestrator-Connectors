package data

import (
    "encoding/xml"
)

// northbound struct
type Northbound_request struct {
	ActionID         string           `json:"ActionID"`
	Target           string           `json:"Target"`
	ActionDefinition Actiondefinition `json:"ActionDefinition"`
}

type Actiondefinition struct {
	ActionType string `json:"ActionType"`
	Service    string `json:"Service"`
	Action     ActionRate `json:"Action"`
}
type ActionRate struct {
	Rate_value int `json:"Rate"`
	Zone string `json:"Zone"`
}

type Northbound_response struct {
	ActionID string `json:"Actionid"`
	Status   string `json:"Status"`
}

///////////////////// southbound struct /////////////////
/*
	type rawXML struct {
		Inner []byte `xml:",innerxml"`
	}

	type Southbound_request struct {
		ITResourceOrchestration ITResourceOrchestration `xml:"ITResourceOrchestration"`
	}

	type ITResourceOrchestration struct {
		ITResource	ITResource `xml:"ITResource"`
	}

	type ITResource struct {
		XMLconfiguration xmlconfiguration `xml:"configuration"`
	}

	type xmlconfiguration struct {
		capability	rawXML `xml:"capability"`
		ConfigurationRule	ConfigurationRule `xml:"configurationRule"`
	}

	type ConfigurationRule struct {
		configurationRuleAction		rawXML `xml:"configurationRuleAction"`
		ConfigurationCondition	ConfigurationCondition `xml:"configurationCondition"`
		externalData		rawXML `xml:"externalData"`
	}

	type ConfigurationCondition struct {
		isCNF	rawXML `xml:"isCNF"`
		DnsRateParameters	DnsRateParameters `xml:"dnsRateParameters"`
	}

	type DnsRateParameters struct {
		operation	string `xml:"operation"`
		IP	string `xml:"ip"`
		Rate	int `xml:"rate"`
	}
*/

type ITResourceOrchestration struct {
    XMLName    xml.Name    `xml:"ITResourceOrchestration"`
    ITResource ITResource  `xml:"ITResource"`
}

type ITResource struct {
    XMLName          xml.Name     `xml:"ITResource"`
    ID               string       `xml:"id,attr"`
    OrchestrationID  string       `xml:"orchestrationID,attr"`
    Configuration    Configuration `xml:"configuration"`
}

type Configuration struct {
    XMLName            xml.Name           `xml:"configuration"`
    XsiType            string             `xml:"xsi:type,attr"`
    Capability         Capability         `xml:"capability"`
    ConfigurationRule  ConfigurationRule  `xml:"configurationRule"`
    Name               string             `xml:"Name"`
}

type Capability struct {
    XMLName xml.Name `xml:"capability"`
    Name    string   `xml:"Name"`
}

type ConfigurationRule struct {
    XMLName                    xml.Name                    `xml:"configurationRule"`
    ConfigurationRuleAction    ConfigurationRuleAction    `xml:"configurationRuleAction"`
    ConfigurationCondition     ConfigurationCondition     `xml:"configurationCondition"`
    ExternalData               ExternalData               `xml:"externalData"`
    Name                       string                     `xml:"Name"`
    IsCNF                      bool                       `xml:"isCNF"`
}

type ConfigurationRuleAction struct {
    XMLName       xml.Name `xml:"configurationRuleAction"`
    XsiType       string   `xml:"xsi:type,attr"`
    DnsActionType string   `xml:"dnsActionType"`
}

type ConfigurationCondition struct {
    XMLName           xml.Name           `xml:"configurationCondition"`
    XsiType           string             `xml:"xsi:type,attr"`
    IsCNF             bool               `xml:"isCNF"`
    DnsRateParameters DnsRateParameters  `xml:"dnsRateParameters"`
}

type DnsRateParameters struct {
    XMLName   xml.Name `xml:"dnsRateParameters"`
    Operation string   `xml:"operation"`
    IP        string   `xml:"ip"`
    Rate      int      `xml:"rate"`
}

type ExternalData struct {
    XMLName xml.Name `xml:"externalData"`
    XsiType string   `xml:"xsi:type,attr"`
    Value   int      `xml:"value"`
}



/*
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

*/