
package acronym

import (
  "fmt"
  "strings"
  "unicode"
)

const testVersion = 1

func abbreviate (input string) (abbr string) {

  runes := []rune(input)

  // loop through the string one rune at a time
  for i, r := range runes {

    // first letter gets abbreviated
    if i == 0 {
      abbr = abbr + string(r)
      continue
    }

    pr := runes[i-1]
    if ( unicode.IsLower(pr) && unicode.IsUpper(r) ) {
      abbr = abbr + string(r)
    }

  }

  abbr = strings.ToUpper(abbr)
  fmt.Println(abbr)
  return
}
