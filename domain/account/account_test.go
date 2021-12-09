package account

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func setup() {
	//mux router with added question routes
	m = mux.NewRouter()

	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestGetAccountProfileByID(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/v1/accounts/profile?id=2", nil)
	if err != nil {
		t.Fatal("Creating 'GET /v1/accounts/profile?id=2' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}

func TestAccountRegistrationHTTPHandler(t *testing.T) {
	setup()
	data := url.Values{}
	data.Set("email", "farras@gmail.com")
	data.Set("password", "farras12345")
	data.Set("firstName", "farras")
	data.Set("lastName", "m")
	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/v1/accounts/registration", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal("Creating 'POST /v1/accounts/registration' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}

func TestAccountLogin(t *testing.T) {
	setup()
	data := url.Values{}
	data.Set("email", "farras@gmail.com")
	data.Set("password", "farras12345")
	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/v1/accounts/authentication", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal("Creating 'POST /v1/accounts/authentication' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}
