package server

import (
	"fmt"
	"net/http"
        "league/challenge/backend/matrix"
)

// General handler that given an uploaded csv file, extract the
// matrix, passes the matrix to the provided matrix function and return
// the string representation return by it. If any error arises when
// extracting or processing the csv file, it will be return to the caller
// as the response with an appropriate status code.
//
// Send request with:
//		curl -F 'file=@/path/matrix.csv' ENDPOINT_URL


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
	http.ListenAndServe(":8080", nil)
}