package server

import (
	"bytes"
	"league/challenge/backend/matrix"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing HTTP error codes and messages

func TestReturnErrorCodeOnMissingFile(t *testing.T) {
	req, err := http.NewRequest("POST", "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handler := handleCsvMatrix(matrix.FmtSquareMatrix)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusBadRequest, status)
	}
}

func TestReturnErrorCodeOnInvalidMatrix(t *testing.T) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "test-matrix.csv")
	part.Write([]byte("1,2,3\n4,5,NaN\n6,7,8"))
	writer.Close()

	req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()

	handler := handleCsvMatrix(matrix.FmtSquareMatrix)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusBadRequest, status)
	}

	expected := "Not a valid number: NaN\n"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}
