/*
Write a program which prompts the user to first enter a name, and then enter an
address. Your program should create a map and add the name and address to the
map using the keys name and address, respectively. Your program should use
Marshal() to create a JSON object from the map, and then your program should
print the JSON object.

Submit your source code for the program,
makejson.go.
*/

package main

import (
  "bufio"
  "fmt"
  "os"
  "encoding/json"
)

func main() {

  var m = make(map[string]string)

  fmt.Print("Enter your name: ")
  nameScanner := bufio.NewScanner(os.Stdin)
  nameScanner.Scan()
  m["name"] = nameScanner.Text()

  fmt.Print("Enter your address: ")
  addressScanner := bufio.NewScanner(os.Stdin)
  addressScanner.Scan()
  m["address"] = addressScanner.Text()

  jsonData, _ := json.Marshal(&m)
  fmt.Println(string(jsonData))

}
