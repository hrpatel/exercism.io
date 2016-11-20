// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

// Complete the type definition.  Pick a suitable data type.
type Clock struct {
  hour, minute int
}


func New(hour, minute int) Clock {
  h := (hour + minute / 60) % 24
  m := minute % 60

  if h < 0 {
    h += 24
  }

  if m < 0 {
    m += 60
  }

  clock :=  Clock{h, m}
  return clock
}

func (clock Clock) String() string {
  return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
  // new_min := (clock.minute + minutes) % 60
  return clock
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
