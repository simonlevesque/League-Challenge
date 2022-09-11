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

	if expected != len(result) {
		t.Errorf("Sum is incorrect : expected %v actual %v", len(result), expected)
	}
}

// Support matrix elements multiplication

func TestMultiplyMatrixElements(t *testing.T) {
	expected := "362880"
	result, _ := Multiply(Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	if expected != result {
		t.Errorf("Multiplication is incorrect : expected %s actual %s", expected, result)
	}
}

// The empty matrix should Multiply to one

func TestMultiplyNothing(t *testing.T) {
	expected := "1"
	result, _ := Multiply(Matrix{})

	if expected != result {
		t.Errorf("Multiplication is incorrect : expected %s actual %s", expected, result)
	}
}

// Negative number changes the sign of the product

func TestMultiplyNegatives(t *testing.T) {
	result, _ := Multiply(Matrix{{1, 1}, {-1, 1}})
	if "-1" != result {
		t.Errorf("Multiplication is incorrect : expected -1 actual %s", result)
	}

	result, _ = Multiply(Matrix{{1, 1}, {-1, -1}})
	if "1" != result {
		t.Errorf("Multiplication is incorrect : expected 1 actual %s", result)
	}
}

// Zero anywhere in the matrix will force the product to zero

func TestMultiplyAnyZero(t *testing.T) {
	result, _ := Multiply(Matrix{{25, 25, 100}, {46, 0, 78}, {403, 50, -56}})
	if "0" != result {
		t.Errorf("Multiplication is incorrect : expected 0 actual %s", result)
	}
}

// Support large numbers multiplication

func TestMultiplyLargeNumbers(t *testing.T) {
	expected := 171
	// use int64 max value
	result, _ := Multiply(Matrix{{9223372036854775807, 9223372036854775807, 9223372036854775807},
		{9223372036854775807, 9223372036854775807, 9223372036854775807},
		{9223372036854775807, 9223372036854775807, 9223372036854775807}})

	if expected != len(result) {
		t.Errorf("Multiplication is incorrect : expecting a number with %v digits, actual %v", expected, len(result))
	}
}

// Invert mirrors matrix on its diagonal

func TestInvertedMatrix(t *testing.T) {
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	result, _ := Invert(Matrix{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}})

	if expected != result {
		t.Errorf("Invert is incorrect : expected %s actual %s", expected, result)
	}
}

// Test simple inverses

func TestTrivialInvert(t *testing.T) {
	result, _ := Invert(Matrix{})

	if "" != result {
		t.Errorf("Nothing should be the inverse of nothing: actual %s", result)
	}

	result, _ = Invert(Matrix{{5}})
	if "5\n" != result {
		t.Errorf("A one by one matrix should be its own inverse: actual %s", result)
	}
}
