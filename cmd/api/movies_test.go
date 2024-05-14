package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMovieHandler_ValidInput(t *testing.T) {
	app := newTestApplication()

	payload := []byte(`{"title":"Inception","year":2010,"runtime":148,"genres":["Action","Adventure","Sci-Fi"]}`)

	req, err := http.NewRequest("POST", "/v1/movies", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.createMovieHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := `{"title":"Inception","year":2010,"runtime":148,"genres":["Action","Adventure","Sci-Fi"]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateMovieHandler_InvalidInput(t *testing.T) {
	app := newTestApplication()

	payload := []byte(`{"title":"Inception","year":"invalid","runtime":148,"genres":["Action","Adventure","Sci-Fi"]}`)

	req, err := http.NewRequest("POST", "/v1/movies", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.createMovieHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}
}

func TestShowMovieHandler_ExistingMovie(t *testing.T) {
	app := newTestApplication()

	req, err := http.NewRequest("GET", "/v1/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.showMovieHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := `{"movie":{"id":1,"createdAt":"[TIMESTAMP]","title":"Casablanca","runtime":102,"genres":["drama","romance","war"],"version":1}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestShowMovieHandler_NonExistingMovie(t *testing.T) {
	app := newTestApplication()

	req, err := http.NewRequest("GET", "/v1/movies/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	app.showMovieHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusNotFound)
	}
}

func newTestApplication() *application {
	return &application{}
}
