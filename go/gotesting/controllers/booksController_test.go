package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"gitlab.com/eliotandelon/gotesting/implementation"
)

var api *API

func TestMain(m *testing.M) {
	service := &implementation.BookServiceMock{}
	api = NewApi(service)
	code := m.Run()
	os.Exit(code)
}

func TestListBooks(t *testing.T) {
	testCases := []struct {
		Name             string
		Params           map[string]string
		ExpectStatusCode int
	}{
		{
			Name:             "All books",
			ExpectStatusCode: 200,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.Name, func(test *testing.T) {
			test.Parallel()
			echoFrme := echo.New()
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			writer := httptest.NewRecorder()
			echoContext := echoFrme.NewContext(request, writer)

			err := api.ListBooks(echoContext)
			if err != nil {
				test.Errorf("unexpected error %s", err)
			}

			if writer.Code != testCase.ExpectStatusCode {
				t.Errorf("Expect status code %v, got %v", testCase.ExpectStatusCode, writer.Code)
			}
		})
	}
}
