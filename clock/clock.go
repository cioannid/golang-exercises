package clock

import "fmt"

const (
  testVersion = 4
  oneDay      = 24 * 60
)

type Clock struct {
  minutes int
}

func New(hour, minute int) Clock {
  minutes := (hour * 60 + minute) % oneDay
  if minutes < 0 {
    minutes = oneDay + minutes
  }
  return Clock{minutes}
}

func (c Clock) String() string {
  return fmt.Sprintf("%02d:%02d", c.minutes / 60, c.minutes % 60)
}

func (c Clock) Add(minutes int) Clock {
  return New(0, c.minutes + minutes)
}
