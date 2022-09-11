package matrix

import (
       "fmt"
       "strconv"
       "strings"
)

// type alias to make signature shorter and more readable

type Matrix = [][]int

// function type that accpts a matric and return a string

type FnMatrixString func(Matrix) (string, error)

// Return a squated formated string representation of a matrix

func FmtSquareMatrix(m Matrix) (string, error) {
     var fmtMatrixStr string
     for _, row := range m {
         var fmtRowStrArray []string
         for _, cell := range row {
             fmtRowStrArray = append(fmtRowStrArray, strconv.Itoa(cell))
         }
         fmtMatrixStr = fmt.Sprintf("%s%s\n", fmtMatrixStr, strings.Join(fmtRowStrArray, ","))
     }
     return fmtMatrixStr, nil
}