package main

import (
	"encoding/json"
	"github.com/Luthfiansyah/warpin-message/app/types"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {

	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetAllMessage(t *testing.T) {
	// Build our expected body

	// expectation
	responseMessage := "Success"

	// Grab our router
	router = gin.Default()
	initRoutes(viper.GetBool("DEBUG_MODE"))

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/v1/message", "") // Assert we encoded correctly,

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code) // Convert the JSON response to a map

	//var response map[string]string
	data := types.GetMessageResponseTest{}

	//fmt.Println(w.Body.String())
	err := json.Unmarshal([]byte(w.Body.String()), &data) // Grab the value & whether or not it exists

	assert.Nil(t, err)
	assert.True(t, data.GeneralResponse.ResponseStatus)
	assert.Equal(t, responseMessage, data.GeneralResponse.ResponseMessage)
}

func TestAddMessage(t *testing.T) {
	// Build our expected body

	// expectation
	responseMessage := "Success"

	// Grab our router
	router = gin.Default()
	initRoutes(viper.GetBool("DEBUG_MODE"))

	body := `{
		"text": "test add message"
	}`

	// Perform a GET request with that handler.
	w := performRequest(router, "POST", "/v1/message", body) // Assert we encoded correctly,

	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code) // Convert the JSON response to a map

	//var response map[string]string
	data := types.AddMessageResponseTest{}

	//fmt.Println(w.Body.String())
	err := json.Unmarshal([]byte(w.Body.String()), &data) // Grab the value & whether or not it exists

	assert.Nil(t, err)
	assert.True(t, data.GeneralResponse.ResponseStatus)
	assert.Equal(t, responseMessage, data.GeneralResponse.ResponseMessage)
}
