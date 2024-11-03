// Echo1 выводит аргументы командной строки
//
// Упражнение 1.1
//
// Измените программу ```echo``` так, чтобы она выводила также os.Args[0], имя выполняемой программы. 


package main

import(
  "fmt"
  "os"
)

func main(){
  var s, sep string
  for i := 0; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "   
  }
  fmt.Println(s)    
}
