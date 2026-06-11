package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VelVit24/todo-api/models"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	user := models.User{Email: "john@doe.com", Password: "password"}
	js, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	w.Write(js)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
