
package hello

const testVersion = 2

// Return a generic or custom hello message.
func HelloWorld(name string) string {
  return_str := "Hello, World!"
  if name != "" {
    return_str = "Hello, " + name + "!"
  }

  return return_str
}
