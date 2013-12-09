package stringer

import "testing"

func AssertEquals(expected, actual interface{}, t *testing.T) {
  if expected != actual {
    t.Errorf("Expected: %v, Actual: %v", expected, actual)
  }
}
