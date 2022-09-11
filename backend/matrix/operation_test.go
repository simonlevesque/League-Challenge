package matrix

import "testing"

func TestSumUpMatrix(t *testing.T) {
     expected := "45"
     result, _ := Sum(Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

     if expected != result {
        t.Errorf("Sum is incorrect : expected %s actual %s", result, expected)
     }
}

// the empty matrix should sum to zero

func TestSumUpNothing(t *testing.T) {
     expected := "0"
     result, _ := Sum(Matrix{})

     if expected != result {
        t.Errorf("Sum is incorrect : expected %s actual %s", result, expected)
     }
}

// Negative number are subtracted from the sum

func TestSubtractNegatives(t *testing.T) {
     expected := "0"
     result, _ := Sum(Matrix{{1, 2}, {0, -3}})

     if expected != result {
        t.Errorf("Sum is incorrect : expected %s actual %s", result, expected)
     }
}

// Support large numbers summation

func TestSumLargeNumbers(t *testing.T) {
     expected := 20
     // use int64 max value
     result, _ := Sum(Matrix{{9223372036854775807, 9223372036854775807, 9223372036854775807},
                             {9223372036854775807, 9223372036854775807, 9223372036854775807},
                             {9223372036854775807, 9223372036854775807, 9223372036854775807}})

     if  expected != len(result) {
        t.Errorf("Sum is incorrect : expected %v actual %v", len(result), expected)
     }
}