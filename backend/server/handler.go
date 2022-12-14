package server

import (
	"fmt"
	"league/challenge/backend/matrix"
	"net/http"
)

// General handler that given an uploaded csv file, extract the
// matrix, passes the matrix to the provided matrix function and return
// the string representation return by it. If any error arises when
// extracting or processing the csv file, it will be return to the caller
// as the response with an appropriate status code.
//
// Send request with:
//		curl -F 'file=@/path/matrix.csv' ENDPOINT_URL
//
// Currently, the handler accepts all HTTP methods, but it could be
// limited to only accept POST or PUT even if they are not a perfect fit
// neither

func handleCsvMatrix(matrixFn matrix.FnMatrixString) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer file.Close()
		matrix, err := matrix.NewMatrix(file)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		response, err := matrixFn(matrix)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Fprint(w, response)
	})
}

// Start the web server lisenning on port 8080

func Start() {
	http.Handle("/echo", handleCsvMatrix(matrix.FmtSquareMatrix))
	http.Handle("/flatten", handleCsvMatrix(matrix.FmtFlattenMatrix))
	http.Handle("/sum", handleCsvMatrix(matrix.Sum))
	http.Handle("/multiply", handleCsvMatrix(matrix.Multiply))
	http.Handle("/invert", handleCsvMatrix(matrix.Invert))
	http.ListenAndServe(":8080", nil)
}
