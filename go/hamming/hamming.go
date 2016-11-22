package hamming

import (
  "errors"
)

const testVersion = 5

func Distance(a, b string) (dist int, err error) {
  // initialize constants
  dist = 0

  // error out on strings with different lengthts
  if len(a) != len(b) {
    err = errors.New("string lengthts dont match!")
    return
  }

	// iterate through each array
  for idx := range a {
    if a[idx] != b[idx] {
      dist++
    }
  }

  return
}
