package matrix

import (
       "fmt"
       "strconv"
       "strings"
)

// type alias to make signature shorter and more readable

type Matrix = [][]int

// Return a squated formated string representation of a matrix

func FmtSquareMatrix(m Matrix) string{
     var fmtMatrixStr string
     for _, row := range m {
         var fmtRowStrArray []string
         for _, cell := range row {
             fmtRowStrArray = append(fmtRowStrArray, strconv.Itoa(cell))
         }
         fmtMatrixStr = fmt.Sprintf("%s%s\n", fmtMatrixStr, strings.Join(fmtRowStrArray, ","))
     }
     return fmtMatrixStr
}