/**
El primer paso poder introducir la ruta.
Segundo paso, añadir git validation, necesito tener controlado esto
*/
package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "strings"
  "io/ioutil"
)
var (
    newFile *os.File
    err     error

)
const inputdelimiter = '\n'
const filename = `querys.txt`

func main() {
  fmt.Print(`Introduce la ruta que quieres leer:`)
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}
  input = strings.Replace(input, "\n", "", -1)
  fmt.Println(input)
  files, _ := ioutil.ReadDir(input)
  for _, f := range files {
    fmt.Println(f.Name())
  }
  Guardar(input)
}
func Guardar(input string){
  fmt.Print(`Guardar el resultado :
    1: Claroooo
    2: nooooo
    Introduce 1 o 2 y pulsa ENTER: `)
  reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}
  switch result {
    case '1':
      fmt.Print(`Guardardando`)
       NewFile(input)
    break
  case '2':
      fmt.Print(`No lo guardas, tu mismo, joder es mi basic de entonces pero noooo`)
    break
  }
}
func NewFile(input string){
  if FileExists() {
    Write(input)
    return
  }
  newFile, err = os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }

    newFile.Close()
    Write(input)
    //FileInfo()
}
func Write(text string){
  f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
      panic(err)
  }

  defer f.Close()

  if _, err = f.WriteString(text+"\n"); err != nil {
      panic(err)
  }
}
func FileInfo(){
  fileInfo, err := os.Stat(filename)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())
    fmt.Println("Is Directory: ", fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n", fileInfo.Sys())
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

}
func FileExists() bool{
  fileInfo, err := os.Stat(filename)
    if err != nil {
        if os.IsNotExist(err) {
            //log.Fatal("File does not exist.")
            return false
        }
    }
    log.Println("File does exist. File information:")
    log.Println(fileInfo)
    return true
}
