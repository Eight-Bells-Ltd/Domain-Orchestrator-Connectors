package logic

import (
	"log"
	"bytes"
	"fmt"
	"io/ioutil"
	//"os"
	"net/http"
	"encoding/xml"
	"github.com/gin-gonic/gin"

	"doc/src/data"
	"doc/src/config"
)

var actionList []data.Northbound_response

func Run() {
	router := gin.Default()

	//add handler function postAction
	router.POST("/EnforceAction", postAction)
	//add handler function getActionList
	router.GET("/getActionList", getActionList)
	//Start REST API server
	//ToDo Set only ePEM IP for security purpose
	router.Run("0.0.0.0:8080")
}

func getActionList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, actionList)
}

func postAction(c *gin.Context) {
	var request data.Northbound_request

	log.Println("EnforceAction Received")

	err := c.BindJSON(&request)

	if err == nil {
		actionstatus := EnforceAction(request)
		if actionstatus {
			response := data.Northbound_response{request.ActionID, "Executed"}
			//Send Response
			c.IndentedJSON(http.StatusOK, response)

			//Update ActionList
			actionList = append (actionList, response)
		} else {
			//ToDo handler enforces action error
			log.Println("Error during Enforce Action")
			response := data.Northbound_response{"ErrorDuringEnforce", "NotExecuted"}
			c.IndentedJSON(http.StatusInternalServerError, response)
			actionList = append (actionList, response)
		}
	} else {
		//Handler parse error
		log.Println("Bad Request=%s", err)
		response := data.Northbound_response{"BadRequest", "NotExecuted"}
		c.IndentedJSON(http.StatusBadRequest, response)
	}

}

func EnforceAction(request data.Northbound_request) bool {
	if request.ActionDefinition.ActionType == "DNS_RATE_LIMIT" {
		//Read File as a template
		templateXml, err := ioutil.ReadFile("templates/dns_rate.xml")
		if err != nil {
			log.Println("ReadFileError=%s", err)
			return false
		}

		fmt.Printf("TEMPLATE: %s\n", templateXml)


		//var Xmltemplate data.Southbound_request
		var Xmltemplate data.ITResourceOrchestration
		err = xml.Unmarshal(templateXml, &Xmltemplate)
		if err != nil {
			log.Println("UnmarshalError=%s", err)
			return false
		}

		Xmltemplate.ITResource.Configuration.ConfigurationRule.ConfigurationCondition.DnsRateParameters.Rate = request.ActionDefinition.Action.Rate_value
		Xmltemplate.ITResource.Configuration.ConfigurationRule.ConfigurationCondition.DnsRateParameters.IP =  "0.0.0.0"
		//xsi:type
		Xmltemplate.ITResource.Configuration.XsiType = "RuleSetConfiguration"
		Xmltemplate.ITResource.Configuration.ConfigurationRule.ConfigurationRuleAction.XsiType = "DNSACTION"
		Xmltemplate.ITResource.Configuration.ConfigurationRule.ConfigurationCondition.XsiType = "DNSCondition"
		Xmltemplate.ITResource.Configuration.ConfigurationRule.ExternalData.XsiType = "Priority"

		XmlRequest, err := xml.MarshalIndent(&Xmltemplate, "", "	")
		if err != nil {
			log.Fatal("MarshalIndentError=%s", err)
			return false
		}

		fmt.Printf("%s\n", XmlRequest)

		//Create Request
		url := config.GetURL()

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(XmlRequest))
		if err != nil {
			log.Println("Error creating request:", err)
			return false
		}

		// Establecer los encabezados
		req.Header.Set("Content-Type", "application/xml")

		//Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending request:", err)
			return false
		}
		defer resp.Body.Close()

		//fmt.Println("Response Status:", resp.Status)
		//fmt.Println("Response Headers:", resp.Header)
		return true
	} else {
		//Handler parse error
		log.Println("ActionType Not supported=%s", request.ActionDefinition.ActionType)
		return false
	}
	
	
}