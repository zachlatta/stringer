package stringer

import "testing"

func TestMin(t *testing.T) {
  AssertEquals(4, min(4, 1337, 12, 1231), t)
}
