
package acronym

import (
  "strings"
  "unicode"
)

const testVersion = 1

func abbreviate (input string) (abbr string) {
  // initialize an array of runes
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
    // handle all caps
    if ( (unicode.IsLower(pr) && unicode.IsUpper(r)) ||
         // handle punctuation and spaces
         (unicode.IsLetter(r) &&
          (unicode.IsSpace(pr) || unicode.IsPunct(pr)))) {

      // add to the abbreviation
      abbr = abbr + string(r)
    }
  }

  abbr = strings.ToUpper(abbr)
  return
}
