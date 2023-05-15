package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("hello world", string(data))
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed ", res.Code)
	// }
}

func TestNamePathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/name", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// nameHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed ", res.Code)
	// }
}

func TestNamePathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/name?name=HY", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// nameHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello HY!", string(data))
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed ", res.Code)
	// }
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"HY","last_name":"Go","email":"abc@dfg.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// assert.Equal(http.StatusOK, res.Code)
	// assert.Equal(http.StatusBadRequest, res.Code)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("HY", user.FirstName)
	assert.Equal("Go", user.LastName)
}
