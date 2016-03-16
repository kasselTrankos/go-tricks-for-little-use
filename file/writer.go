package file

import(
  "log"
  "os"
)


func Save(input string){
  if FileExists() {
    Write(input)
    return
  }
  newFile, err = os.Create(Filename)
    if err != nil {
        log.Fatal(err)
    }
    newFile.Close()
    Write(input)
}


func Write(text string){
  f, err := os.OpenFile(Filename, os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
      panic(err)
  }
  defer f.Close()
  if _, err = f.WriteString(text+"\n"); err != nil {
      panic(err)
  }
}
func FileExists() bool{
  fileInfo, err := os.Stat(Filename)
    if err != nil {
        if os.IsNotExist(err) {
            //log.Fatal("File does not exist.")
            return false
        }
    }
    log.Println("File does exist. File information:\n")
    log.Println(fileInfo)
    return true
}
