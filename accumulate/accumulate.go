package accumulate

const testVersion = 1

func Accumulate(in []string, f func(string) string) (out []string) {
  for _, v := range in {
    out = append(out, f(v))
  }
  return
}
