package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"student-api/database"
)

func resetData() {
	database.Students = nil
	database.NextID = 1
}

func TestHealthCheck(t *testing.T) {

	router := SetupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/healthcheck",
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestCreateStudent(t *testing.T) {

	resetData()

	router := SetupRouter()

	body := []byte(`{
		"name":"Mayur",
		"email":"mayur@test.com"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/students",
		bytes.NewBuffer(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 got %d", w.Code)
	}
}

func TestGetStudents(t *testing.T) {

	resetData()

	router := SetupRouter()

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/students",
		bytes.NewBuffer([]byte(`{
			"name":"Mayur",
			"email":"mayur@test.com"
		}`)),
	)

	createReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	router.ServeHTTP(
		httptest.NewRecorder(),
		createReq,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/students",
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestGetStudent(t *testing.T) {

	resetData()

	router := SetupRouter()

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/students",
		bytes.NewBuffer([]byte(`{
			"name":"Mayur",
			"email":"mayur@test.com"
		}`)),
	)

	createReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	router.ServeHTTP(
		httptest.NewRecorder(),
		createReq,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/students/1",
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestUpdateStudent(t *testing.T) {

	resetData()

	router := SetupRouter()

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/students",
		bytes.NewBuffer([]byte(`{
			"name":"Mayur",
			"email":"mayur@test.com"
		}`)),
	)

	createReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	router.ServeHTTP(
		httptest.NewRecorder(),
		createReq,
	)

	updateReq := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/students/1",
		bytes.NewBuffer([]byte(`{
			"name":"Mayur Kadam",
			"email":"mayur.kadam@test.com"
		}`)),
	)

	updateReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, updateReq)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestDeleteStudent(t *testing.T) {

	resetData()

	router := SetupRouter()

	createReq := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/students",
		bytes.NewBuffer([]byte(`{
			"name":"Mayur",
			"email":"mayur@test.com"
		}`)),
	)

	createReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	router.ServeHTTP(
		httptest.NewRecorder(),
		createReq,
	)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/students/1",
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected 204 got %d", w.Code)
	}
}