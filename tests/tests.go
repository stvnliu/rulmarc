package tests

import "fmt"

func DoThis(str string) string {
  fmt.Println(str)
  return fmt.Sprintf("This is the string returned: %v", str)
}
