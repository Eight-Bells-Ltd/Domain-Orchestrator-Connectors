package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var actionList []data.northbound_response

func run() {
	router := gin.Default()

	//add handler function postAction
	router.POST("/EnforceAction", postAction)

	router.GET("/getActionList", getActionList)
	//Start REST API server
	router.Run("localhost:8080")
}

func getActionList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, actionList)
}

func postAction(c *gin.Context) {
	var request data.northbound_request

	err := c.BindJSON(&request)

	if err == nil {
		actionstatus := EnforceAction(request)
		if actionstatus {
			response := data.northbound_response{request.actionID, "Executed"}
			c.IndentedJSON(http.StatusOK, response)
			actionList.add(response)
		} else {
			//ToDo handler enforces action error
		}
	} else {
		//Handler parse error
		response := data.northbound_response{"BadRequest", "NotExecuted"}
		c.IndentedJSON(http.StatusBadRequest, response)
		actionList.add(response)
	}

}

func EnforceAction(request data.northbound_request) bool {
	//ToDo Define how to enforce action
	return true
}
