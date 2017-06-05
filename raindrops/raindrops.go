package raindrops

import "strconv"

const testVersion = 3

var factors = []struct {
  number  int
  phrase  string
}{
  {3, "Pling"},
  {5, "Plang"},
  {7, "Plong"},
}

func Convert(x int) (result string) {
  hasFactor := false

  for _, factor := range factors {
    if x % factor.number == 0 {
      hasFactor = true
      result = result + factor.phrase
    }
  }
  if !hasFactor {
    result = strconv.Itoa(x)
  }
  return
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
