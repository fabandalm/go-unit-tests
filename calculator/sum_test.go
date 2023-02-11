package calculator_test

import (
	"go_unit_testing/calculator"
	"testing"
)

type TestCase struct {
	num_a    int
	numb_b   int
	expected int
	actual   int
}

func TestSumTwoNumbers_PositiveNumbers(t *testing.T) {
	testCase := TestCase{
		num_a:    2,
		numb_b:   3,
		expected: 5,
	}

	testCase.actual = calculator.SumTwoNumbers(testCase.num_a, testCase.numb_b)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestSumTwoNumbers_NegativeNumbers(t *testing.T) {
	testCase := TestCase{
		num_a:    -5,
		numb_b:   -2,
		expected: -7,
	}

	testCase.actual = calculator.SumTwoNumbers(testCase.num_a, testCase.numb_b)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
