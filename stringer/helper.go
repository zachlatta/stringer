package stringer

func min(args ...int) int {
  min := args[0]
  for _, v := range args {
    if min > v {
      min = v
    }
  }
  return min
}
