package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(10, 5)
	assert.Equal(t, result, 15)
}

func TestDiff(t *testing.T) {
	result := Diff(10, 5)
	assert.Equal(t, result, 5)
}

func TestProduct(t *testing.T) {
	result := Product(10, 5)
	assert.Equal(t, result, 50)
}

func TestDivide(t *testing.T) {
	result := Divide(10, 5)
	assert.Equal(t, result, 2)
}

func TestRemainder(t *testing.T) {
	result := Remainder(13, 5)
	assert.Equal(t, result, 3)
}

func TestFibonacci(t *testing.T) {
	var result []int = Fibonacci(0, 10)
	assert.Equal(t, result, []int{0, 1, 1, 2, 3, 5, 8})
}

func TestAddFibonacci(t *testing.T) {
	result := AddFibonacci(1, 10)
	assert.Equal(t, result, 20)
}
