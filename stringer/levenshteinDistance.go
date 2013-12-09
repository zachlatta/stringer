package stringer

func LevenshteinDistance(s string, t string) int {
  // Degenerate cases
  if s == t {
    return 0
  }
  if len(s) == 0 {
    return len(t)
  }
  if len(t) == 0 {
    return len(s)
  }

  // Create two work vectors of integer distances
  v0 := make([]int, len(t) + 1)
  v1 := make([]int, len(t) + 1)

  // Initialize v0 (previous row of distances)
  // this row is A[0][i] - edit distance 
  // the distance is just the number of characters to delete from t
  for i := range v0 {
    v0[i] = i;
  }

  // Calculate v1 (current row distances) from the previous row v0
  for i := range s {
    // First element of v1 is A[i+1][0]
    // Edit distance is delete (i+1) chars from s to match empty t
    v1[0] = i + 1

    // Use formula to fill in the rest of the row
    for j := range t {
      var cost int
      if s[i] == t[j] {
        cost = 0
      } else {
        cost = 1
      }

      v1[j + 1] = min(v1[j] + 1, v0[j + 1] + 1, v0[j] + cost)
    }

    // Copy v1 (current row) to v0 (previous row) for next iteration
    for j := range v0 {
      v0[j] = v1[j]
    }
  }

  return v1[len(t)]
}
