package gigasecond

// import path for the time package from the standard library
import "time"

const (
  testVersion = 4
  gigaSecond  = 1e9
)

func AddGigasecond(t time.Time) time.Time {
  return t.Add(time.Duration(gigaSecond) * time.Second)
}
