package matrix

import (
       "encoding/csv"
       "errors"
       "fmt"
       "io"
       "strconv"
       "strings"
)

// Given a csv file, return the matrix representation or an error

func NewMatrix(r io.Reader) (Matrix, error) {
     records, err := csv.NewReader(r).ReadAll()
     if err != nil {
	    return nil, fmt.Errorf("Invalid CSV file: %s", err.Error())
     }
     dimension := len(records)

     matrix := make(Matrix, dimension)

     for i, row := range records {

         // If at lease one of the row size is not equal to the number
         // of columns, then it cannot be a square matrix
         if len(row) != dimension {
            return nil, errors.New("Matrix is not square")
         }
         matrix[i] = make([]int, dimension)
         for j, elm := range row {
             value, err := strconv.Atoi(strings.TrimSpace(elm))

             if err != nil {
                return nil, errors.New(fmt.Sprintf("Not a valid number: %s", elm))
             }

             matrix[i][j] = value
         }
     }

     return matrix, nil
}