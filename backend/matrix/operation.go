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