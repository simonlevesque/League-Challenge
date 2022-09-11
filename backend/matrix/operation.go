package matrix

import "math/big"

func Sum(m Matrix) (string, error) {
     s := big.NewInt(0)
     for _, row := range m {
         for _, cell := range row {
             s.Add(s, big.NewInt(int64(cell)))
         }
     }
     return s.Text(10), nil
}

func Multiply(m Matrix) (string, error) {
     p := big.NewInt(1)
     for _, row := range m {
         for _, cell := range row {
             p.Mul(p, big.NewInt(int64(cell)))
         }
     }
     return p.Text(10), nil
}

func Invert(m Matrix) (string, error) {
     dimension := len(m)
     mirror := make(Matrix, dimension)
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