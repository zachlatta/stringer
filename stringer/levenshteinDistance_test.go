package stringer

import "testing"

func TestLevenshteinDistance(t *testing.T) {
  expected := 19
  actual := LevenshteinDistance("these", "strings do not match!")

  AssertEquals(expected, actual, t)
}
