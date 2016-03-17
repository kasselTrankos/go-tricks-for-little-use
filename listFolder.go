/**
El primer paso poder introducir la ruta.
Segundo paso, añadir git validation, necesito tener controlado esto
*/
package main

import (
    "./file"
    "./snif"
  "bufio"
  "fmt"
  "os"
  "strings"
  "io/ioutil"
)

const inputdelimiter = '\n'
func main(){
  fmt.Print(`Que hacer :
    1: Analysis
    2: File

    Introduce 1 o 2 y pulsa ENTER: `)
  reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}
  switch result {
  case '1':
    snif.Sniff()
    break
  case '2':
      Route()
    break
  }
}

func Route() {
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
  OutResponse(input)
}

func OutResponse(input string){
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
      fmt.Print("Guardardando....\n")
       file.Save(input)
    break
  case '2':
      fmt.Print(`No lo guardas, tu mismo, joder es mi basic de entonces pero noooo`)
    break
  }
}
