package basics

import(
  "net/http"
  "io/ioutil"
  "fmt"
)
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/4037124?start=0
func BlankIdentifierExample(){
  res, _ := http.Get("https://blockexplorer.com/api/addr/1Q85UwrAUKm4EXE5jSmpoyTiys8BCos45J")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", page)
}
