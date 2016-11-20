
// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock data type.
type Clock struct {
  hour, minute int
}


func New(hour, minute int) Clock {
  h := (hour + minute / 60) % 24
  m := minute % 60

  if m < 0 {
    m += 60
    h -= 1
  }

  if h < 0 {
    h += 24
  }

  return Clock{h, m}
}

func (clock Clock) String() string {
  return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
  return New(clock.hour, clock.minute + minutes)
}
