package matrix

// Collection of related matrix transformation excluding formatting, printing or parsing

import "math/big"

// Given a matrix, add up all its values from left to right, top to bottom
// Big integers are used here to avoid overflows with large
// matrixes. Note that at the moment of writing, matrixes were
// implemented as integer and would not overflow from a cell by itself

func Sum(m Matrix) (string, error) {
     s := big.NewInt(0)
     for _, row := range m {
         for _, cell := range row {
             s.Add(s, big.NewInt(int64(cell)))
         }
     }
     return s.Text(10), nil
}

// Given a matrix, multiply all its value together
// Big integers are used here to avoid overflows with large
// matrixes. Note that at the moment of writing, matrixes were
// implemented as integer and would not overflow from a cell by itself

func Multiply(m Matrix) (string, error) {
     p := big.NewInt(1)
     for _, row := range m {
         for _, cell := range row {
             p.Mul(p, big.NewInt(int64(cell)))
         }
     }
     return p.Text(10), nil
}

// Given a matrix, mirror it on its diagonal left to right, top to bottom

func Invert(m Matrix) (string, error) {
     dimension := len(m)
     mirror := make(Matrix, dimension)

     // Allocate row memory before using it
     for i, _ := range m {
         mirror[i] = make([]int, dimension)
     }
     
     for i, row := range m {
         for j, cell := range row {
             mirror[j][i] = cell
         }
     }
     return FmtSquareMatrix(mirror)
}