package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sarikap08/restProject/pkg/handlers"
	"github.com/stretchr/testify/suite"
)

type SystemHandlerTestSuite struct {
	suite.Suite
	systemHandler *handlers.SystemHandler
}

func (suite *SystemHandlerTestSuite) SetupSuite() {
	suite.systemHandler = handlers.NewSystemHandler()
}

func (suite *SystemHandlerTestSuite) TestGetMachine() {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5051/machines", nil)

	response := executeRequest(request)

	suite.systemHandler.GetMachine(response, request)

	// b, _ := io.ReadAll(response.Body)
	// if body := string(b); body != "[]" {
	// 	suite.ErrorAsf()
	// 	suite.ErrorAsf("Expected an empty array. Got %s", body)
	// }

	suite.Equal(response.Result().StatusCode, http.StatusOK)

}

func (suite *SystemHandlerTestSuite) TestCreateMachine() {
	var jsonStr = []byte(`{
		"machineId": 12345,
		"stats": {
			"cpuTemp": 90,
			"fanSpeed": 401,
			"HDDSpace": 801
		},
		"lastLoggedIn": "admin/xxx",
		"sysTime": "2022-04-23T18:25:43.511Z"
	}`)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:5051/machines", bytes.NewBuffer(jsonStr))

	resp := executeRequest(req)

	suite.systemHandler.CreateMachine(resp, req)

	suite.Equal(resp.Result().StatusCode, http.StatusCreated)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	return rr
}

func TestSystemHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SystemHandlerTestSuite))
}
