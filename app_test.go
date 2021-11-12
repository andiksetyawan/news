package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"news/model"
	"testing"
	"time"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPostTagRouteSuccess(t *testing.T) {
	router := setupRouter()

	request := model.TagCreateRequest{Name: "Tag Test " + fmt.Sprint(time.Now().Unix())}
	requestByte, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestByte)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tag", requestReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPostTagRouteError(t *testing.T) {
	router := setupRouter()

	request := model.TagCreateRequest{Name: ""}
	requestByte, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestByte)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tag", requestReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestGetTagRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tag/investasi", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response bson.M
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, response["success"], true)
	assert.Equal(t, response["data"].(map[string]interface{})["name"], "Investasi")
}

func TestGetsTagRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tags", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response bson.M
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, response["success"], true)
	assert.Greater(t, len(response["data"].([]interface{})), 0)
}

func TestPutTagRouteSuccess(t *testing.T) {
	router := setupRouter()

	request := model.TagCreateRequest{Name: "Tag Test Update" + fmt.Sprint(time.Now().Unix())}
	requestByte, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestByte)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/tag/1", requestReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPutTagRouteError(t *testing.T) {
	router := setupRouter()

	request := model.TagCreateRequest{Name: "Tag Test Update Error" + fmt.Sprint(time.Now().Unix())}
	requestByte, _ := json.Marshal(request)
	requestReader := bytes.NewReader(requestByte)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/tag/00000xxx", requestReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
