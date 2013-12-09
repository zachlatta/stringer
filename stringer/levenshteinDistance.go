package stringer

func LevenshteinDistance(s string, t string) int {
  return levenshteinDistance(s, len(s), t, len(t))
}

func levenshteinDistance(s string, lenS int, t string, lenT int) int {
  if lenS == 0 {
    return lenT
  }
  if lenT == 0 {
    return lenS
  }

  var cost int

  if s[lenS-1] == t[lenT-1] {
    cost = 0
  } else {
    cost = 1
  }

  return min(levenshteinDistance(s, lenS - 1, t, lenT) + 1,
    levenshteinDistance(s, lenS, t, lenT - 1) + 1,
    levenshteinDistance(s, lenS - 1, t, lenT - 1) + cost)
}
