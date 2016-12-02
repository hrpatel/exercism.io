
package acronym

import (
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

    // previous rune/letter
    pr := runes[i-1]
    if ( (unicode.IsLower(pr) && unicode.IsUpper(r)) ||
         (unicode.IsLetter(r) &&
          (unicode.IsSpace(pr) || unicode.IsPunct(pr)))) {
      abbr = abbr + string(r)
    }

  }

  abbr = strings.ToUpper(abbr)
  return
}
