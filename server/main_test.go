package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNumbersAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRoutes()
	
	tests := []struct {
		name          string
		initNumbers   []int
		number        int
		expectedCode  int
		expectedError string
		expectedList  []int
	}{
		{
			name:         "Add first positive number",
			number:       5,
			expectedCode: http.StatusOK,
			expectedList: []int{5},
		},
		{
			name:          "Reject zero",
			number:        0,
			expectedCode:  http.StatusBadRequest,
			expectedError: "Zero is not allowed",
		},
		{
			name:         "Add first negative number",
			number:       -3,
			expectedCode: http.StatusOK,
			expectedList: []int{-3},
		},
		{
			name:         "Add second positive number",
			initNumbers:  []int{5},
			number:       3,
			expectedCode: http.StatusOK,
			expectedList: []int{5, 3},
		},
		{
			name:         "Add second negative number",
			initNumbers:  []int{-3},
			number:       -2,
			expectedCode: http.StatusOK,
			expectedList: []int{-3, -2},
		},
		{
			name:         "Subtract from [40,40] yields -1",
			initNumbers:  []int{40, 40},
			number:       -81,
			expectedCode: http.StatusOK,
			expectedList: []int{-1},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers = []int{}
			if len(tt.initNumbers) > 0 {
				numbers = tt.initNumbers
			}
			
			reqBody := map[string]int{"number": tt.number}
			jsonBody, _ := json.Marshal(reqBody)
			
			req, _ := http.NewRequest("POST", "/api/numbers", bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			
			assert.Equal(t, tt.expectedCode, w.Code)
			
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			
			if tt.expectedError != "" {
				assert.Equal(t, tt.expectedError, response["error"])
			} else {
				list := response["list"].([]interface{})
				var resultList []int
				for _, v := range list {
					resultList = append(resultList, int(v.(float64)))
				}
				assert.Equal(t, tt.expectedList, resultList)
			}
		})
	}
}

func TestGetNumbers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRoutes() // ...existing code...
	numbers = []int{1, 2, 3} // Set up initial state
	
	req, _ := http.NewRequest("GET", "/api/numbers", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	list := response["list"].([]interface{})
	var resultList []int
	for _, v := range list {
		resultList = append(resultList, int(v.(float64)))
	}
	assert.Equal(t, []int{1, 2, 3}, resultList)
}
