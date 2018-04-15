package basics

//A runes is: a single character, from any Alphabet, from any language, in the world:
//character
//integer
//alias for int32
//how many bytes in 32 bits? (4 bytes -> 4 * 8 = 32)
//UTF-8 is a 4 byte coding scheme
//with int 32 (4 bytes) we have numbers for each of the code points
//string:  "" or ``
//rune: ''


import "fmt"

func Cast(){
  fmt.Println([]byte("Hello, playground"))
}

func RuneCounter(){
  for i := 50; i <= 140; i++{
    fmt.Println('i', " - ", string(i), " - ", []byte(string(i)));
  }
  foo := 'a'
  fmt.Println(foo)
  fmt.Printf("%T \n", foo)
  fmt.Printf("%T \n", rune(foo))
}
