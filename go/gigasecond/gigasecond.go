// Package clause.
package gigasecond

import (
  "time"
)

// Constant declaration.
const testVersion = 4

// API function.  It uses a type from the Go standard library.
func AddGigasecond(curr_time time.Time) time.Time {
  return curr_time.Add(1000000000 * time.Second)
}
