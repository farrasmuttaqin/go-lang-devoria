package article

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

func TestCreateArticle(t *testing.T) {
	setup()
	data := url.Values{}
	data.Set("title", "artikel memasak")
	data.Set("subtitle", "memasak dirumah membuat anda lebih sehat")
	data.Set("content", "memasak dirumah jauh lebih sehat untuk kesehatan anda. Karena anda dapat meracik sendiri sesuai kebutuhan anda terhadap makanan yang anda makan.")
	data.Set("author", "2")
	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/v1/article/create", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal("Creating 'POST /v1/article/create' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}

func TestCreateArticle(t *testing.T) {
	setup()
	data := url.Values{}
	data.Set("title", "artikel memasak")
	data.Set("subtitle", "memasak dirumah membuat anda lebih sehat")
	data.Set("content", "memasak dirumah jauh lebih sehat untuk kesehatan anda. Karena anda dapat meracik sendiri sesuai kebutuhan anda terhadap makanan yang anda makan.")
	data.Set("author", "2")
	//Testing get of non existent question type
	req, err = http.NewRequest("POST", "/v1/article/create", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal("Creating 'POST /v1/article/create' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}

func TestGetArticleByID(t *testing.T) {
	setup()
	//Testing get of non existent question type
	req, err = http.NewRequest("GET", "/v1/article/get?id=9", nil)
	if err != nil {
		t.Fatal("Creating 'GET /v1/article/get?id=9' request failed!")
	}

	m.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusNotFound)
	}
}
