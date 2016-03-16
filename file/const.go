package file
import(
  "os"
)

const Filename string = "querys";
var (
    newFile *os.File
    err     error

)

func Name() string {
  return Filename
}
