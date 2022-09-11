package matrix

import (
	"reflect"
	"strings"
	"testing"
)

// Small helper function to avoid error handling when printing matrixes

func prtMatrix(m Matrix) string {
	s, _ := FmtSquareMatrix(m)
	return s
}

func TestCreatNewMatrix(t *testing.T) {
	csvMockData := "1, 2, 3\n 4, 5, 6\n 7, 8, 9"
	mockCsvFileReader := strings.NewReader(csvMockData)

	expected := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	matrix, err := NewMatrix(mockCsvFileReader)

	if err != nil {
		t.Errorf("Matrix creation failed error: %s", err.Error())
	}

	if !reflect.DeepEqual(expected, matrix) {
		t.Errorf("Matrix creation return wrong matrix: \n\nexpected: \n%s\nactual: \n%s", prtMatrix(expected), prtMatrix(matrix))
	}
}

// Testing that only square matrix are accepted

func TestFailOnNonSquareMetrix(t *testing.T) {
	csvMockData := "1, 2, 3, 4, 5, 6\n 7, 8, 9\n, 10\n 11, 12"
	mockCsvFileReader := strings.NewReader(csvMockData)

	matrix, err := NewMatrix(mockCsvFileReader)

	if matrix != nil {
		t.Errorf("Matrix creation should have failed")
	}

	if err == nil {
		t.Errorf("Error should have been return on matrix creation failure")
	}
}

// Testing that non integer values are rejected

func TestFailOnInvalidNumber(t *testing.T) {
	csvMockData := "1, 2, NaN\n 4, 5, 6\n 7, 8, 9"
	mockCsvFileReader := strings.NewReader(csvMockData)

	matrix, err := NewMatrix(mockCsvFileReader)

	if matrix != nil {
		t.Errorf("Matrix creation should have failed")
	}

	if err == nil {
		t.Errorf("Error should have been return on matrix creation failure")
	}
}

func TestFailOnDecimalNumber(t *testing.T) {
	csvMockData := "1, 2, 3.04\n 4, 5, 6\n 7, 8, 9"
	mockCsvFileReader := strings.NewReader(csvMockData)

	matrix, err := NewMatrix(mockCsvFileReader)

	if matrix != nil {
		t.Errorf("Matrix creation should have failed")
	}

	if err == nil {
		t.Errorf("Error should have been return on matrix creation failure")
	}
}

// Testing that negative values are accepted

func TestNegativeNumberAreAllowed(t *testing.T) {
	csvMockData := "1, 2, -3\n 4, -5, 6\n 7, 8, 9"
	mockCsvFileReader := strings.NewReader(csvMockData)

	expected := Matrix{{1, 2, -3}, {4, -5, 6}, {7, 8, 9}}

	matrix, err := NewMatrix(mockCsvFileReader)

	if err != nil {
		t.Errorf("Matrix creation failed error: %s", err.Error())
	}

	if !reflect.DeepEqual(expected, matrix) {
		t.Errorf("Matrix creation return wrong matrix: \n\nexpected: \n%s\nactual: \n%s", prtMatrix(expected), prtMatrix(matrix))
	}
}

// Testing that empty file returns an empty matrix

func TestEmptyFileReturnEmptyMatrix(t *testing.T) {
	csvMockEmptyData := ""
	mockCsvEmptyFileReader := strings.NewReader(csvMockEmptyData)

	expected := Matrix{}

	emptyMatrix, err := NewMatrix(mockCsvEmptyFileReader)

	if err != nil {
		t.Errorf("Matrix creation failed error: %s", err.Error())
	}

	if !reflect.DeepEqual(expected, emptyMatrix) {
		t.Errorf("Matrix creation return wrong matrix: \n\nexpected: \n%s\nactual: \n%s", prtMatrix(expected), prtMatrix(emptyMatrix))
	}

}
