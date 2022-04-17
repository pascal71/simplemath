/**
 * Author: Pascal van Dam
 * File: simplemath_test.go
 */

package simplemath_test

import (
	"testing"

	"github.com/pascal71/simplemath"
)

func TestAdd(t *testing.T) {
	var a, b int
	var expected int

	a = 4
	b = 3
	expected = 7

	if got := simplemath.Add(a, b); got != expected {
		t.Errorf("Add(%d,%d) = %d, didn't return %d", a, b, got, expected)
	}
}
