package server

import (
        "bytes"
        "league/challenge/backend/matrix"
        "mime/multipart"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestEchoHandler(t *testing.T) {

        body := bytes.NewBuffer(nil)
        writer := multipart.NewWriter(body)

        part, err := writer.CreateFormFile("file", "test-matrix.csv")
        part.Write([]byte("1,2,3\n4,5,6\n7,8,9"))
        writer.Close()

        req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}
        req.Header.Set("Content-Type", writer.FormDataContentType())

        rec := httptest.NewRecorder()

        handler := handleCsvMatrix(matrix.FmtSquareMatrix)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusOK, status)
	}

        expected := "1,2,3\n4,5,6\n7,8,9\n"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}

func TestFlattenHandler(t *testing.T) {

        body := bytes.NewBuffer(nil)
        writer := multipart.NewWriter(body)

        part, err := writer.CreateFormFile("file", "test-matrix.csv")
        part.Write([]byte("1,2,3\n4,5,6\n7,8,9"))
        writer.Close()

        req, err := http.NewRequest("POST", "/flatten", body)
	if err != nil {
		t.Fatal(err)
	}
        req.Header.Set("Content-Type", writer.FormDataContentType())

        rec := httptest.NewRecorder()

        handler := handleCsvMatrix(matrix.FmtFlattenMatrix)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusOK, status)
	}

        expected := "1,2,3,4,5,6,7,8,9"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}

func TestSumHandler(t *testing.T) {

        body := bytes.NewBuffer(nil)
        writer := multipart.NewWriter(body)

        part, err := writer.CreateFormFile("file", "test-matrix.csv")
        part.Write([]byte("1,2,3\n4,5,6\n7,8,9"))
        writer.Close()

        req, err := http.NewRequest("POST", "/sum", body)
	if err != nil {
		t.Fatal(err)
	}
        req.Header.Set("Content-Type", writer.FormDataContentType())

        rec := httptest.NewRecorder()

        handler := handleCsvMatrix(matrix.Sum)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusOK, status)
	}

        expected := "45"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}

func TestMultiplyHandler(t *testing.T) {

        body := bytes.NewBuffer(nil)
        writer := multipart.NewWriter(body)

        part, err := writer.CreateFormFile("file", "test-matrix.csv")
        part.Write([]byte("1,2,3\n4,5,6\n7,8,9"))
        writer.Close()

        req, err := http.NewRequest("POST", "/multiply", body)
	if err != nil {
		t.Fatal(err)
	}
        req.Header.Set("Content-Type", writer.FormDataContentType())

        rec := httptest.NewRecorder()

        handler := handleCsvMatrix(matrix.Multiply)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusOK, status)
	}

        expected := "362880"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}