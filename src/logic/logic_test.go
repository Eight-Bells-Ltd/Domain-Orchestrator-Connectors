package logic

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/require"
    "encoding/json"
    "bytes"

	"doc/src/data"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestNorthBoundError(t *testing.T) {
	router := SetUpRouter()
    router.POST("/EnforceAction", postAction)

	// Create request with empty body and check error
    req, err := http.NewRequest("POST", "/EnforceAction", nil)
    if err != nil {
        t.Fatal(err)
    }

    record := httptest.NewRecorder()

    router.ServeHTTP(record, req)

    //Check response code
    if status := record.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    //Check response body
    expected := `{"actionid": "BadRequest","status": "NotExecuted"}`
    require.JSONEq(t, expected, record.Body.String())
}

func TestNorthBoundSuccess(t *testing.T) {
    router := SetUpRouter()
    router.POST("/EnforceAction", postAction)

	// Create request with correct body
    body := data.Northbound_request {
        ActionID: "value",
        Target: "value",
        ActionDefinition: data.Actiondefinition {
            ActionType: "value",
            Service: "value",
            Action: data.ActionRate {
                Zone: "0.0.0.0",
                Rate_value: 5,
            },
        },
    }
    jsonbody, _ := json.Marshal(body)

    req, err := http.NewRequest("POST", "/EnforceAction", bytes.NewBuffer(jsonbody))
    if err != nil {
        t.Fatal(err)
    }

    record := httptest.NewRecorder()

    router.ServeHTTP(record, req)

    //Check response code
    if status := record.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    //Check response body
    expected := `{"actionid": "value","status": "Executed"}`
    require.JSONEq(t, expected, record.Body.String())
}
