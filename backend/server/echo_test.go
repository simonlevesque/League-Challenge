package server

import (
        "bytes"
        "mime/multipart"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestEchoHandler(t *testing.T) {

        body := bytes.NewBuffer(nil)
        writer := multipart.NewWriter(body)

        part, err := writer.CreateFormFile("file", "test-matrix.csv")
        part.Write([]byte("1, 2, 3\n4, 5, 6\n7, 8, 9"))
        writer.Close()

        req, err := http.NewRequest("POST", "/echo", body)
	if err != nil {
		t.Fatal(err)
	}
        req.Header.Set("Content-Type", writer.FormDataContentType())

        rec := httptest.NewRecorder()

        handler := http.HandlerFunc(Echo)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Echo handler returned wrong status code: expected %v actual %v",
			http.StatusOK, status)
	}

        expected := "1, 2, 3\n4, 5, 6\n7, 8, 9\n"
	if rec.Body.String() != expected {
		t.Errorf("Echo handler returned wrong body: expected %v actual %v",
			expected, rec.Body.String())
	}
}