package logic

import (
	"log"
	//"os"
	"net/http"
	//"encoding/xml"
	"github.com/gin-gonic/gin"

	"doc/src/data"
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

	log.Print("EnforceAction Received")

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
		}
	} else {
		//Handler parse error
		log.Printf("Bad Request=%s", err)
		response := data.Northbound_response{"BadRequest", "NotExecuted"}
		c.IndentedJSON(http.StatusBadRequest, response)
		actionList = append (actionList, response)
	}

}

func EnforceAction(request data.Northbound_request) bool {
	//ToDo Define how to enforce action
	/*
	if request==nil {
		xmlData, err := createXMLBody("rateLimit")
		if err != nil
			return true
	}
	*/

	return true
}

//ToDo create XML body request reading a template
/*
func createXMLBody(actionType string) ([]byte, error) {
    xmlBody := data.Southbound_request{
        XMLName: "ToDo",
    }

	if actionType == "rateLimit" {
		readXMLTemplate("templates/dns_rate.xml")
	}

    return xmlBody, nil
}

func readXMLTemplate(filename string) ([]byte, error) {
    return os.ReadFile(filename)
}
*/