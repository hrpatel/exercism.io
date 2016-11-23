
package acronym

import (
  "fmt"
  "strings"
)

const testVersion = 1

func abbreviate (input string) (abbr string) {

  // Return true if ' ', ',', ':', '-'
  f := func(c rune) bool {
    return c == ',' || c == ':' || c == ' ' || c == '-'
  }

  parts := strings.FieldsFunc(input, f)
  fmt.Println(parts)

  for _, part := range parts {
    fmt.Println(string(part[0]))
    abbr = abbr + strings.ToUpper(string(part[0]))
  }

  return
}
