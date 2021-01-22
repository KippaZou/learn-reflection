package test

import (
	"fmt"
	"testing"
)

func TestMakeSum(t *testing.T) {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	MakeSum(&intSum)
	MakeSum(&floatSum)
	MakeSum(&stringSum)

	fmt.Println(intSum(1, 2))
	fmt.Println(floatSum(2.1, 3.5))
	fmt.Println(stringSum("Geeksfor", "Geeks"))
}
