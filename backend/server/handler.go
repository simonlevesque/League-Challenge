package server

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

// Given an uploaded csv file, return the matrix as a string in matrix format.
//
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func Echo(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	}

// Start the web server lisenning on port 8080

func Start() {
	http.HandleFunc("/echo", Echo)
	http.ListenAndServe(":8080", nil)
}