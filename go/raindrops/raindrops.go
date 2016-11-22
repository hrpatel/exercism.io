package raindrops

import (
  "fmt"
)

const testVersion = 2

func Convert(input int) (sound string) {

  // define the sound mapping
  var sounds = map[int]string {
    3: "Pling",
    5: "Plang",
    7: "Plong",
  }

  // check for factor and add to the sound
  for _, x := range []int{3,5,7} {
    if input % x == 0 {
      sound = sound + sounds[x]
    }
  }

  // no sound was made so lets pass the input through
  if sound == "" {
    sound = fmt.Sprintf("%d", input)
  }

  return
}
