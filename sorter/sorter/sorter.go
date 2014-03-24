package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func read_values(infile string)(values [] int, err error) {
  file, err := os.Open(infile)
  if err != nil {
    fmt.Println("Failed to open file ", infile)
    return
  }
  defer file.Close()

  br := bufio.NewReader(file)
  values = make([] int, 0)

  for {
    line, is_prefix, err1 := br.ReadLine()
    if err1 != nil {
      if err1 != io.EOF {
        err = err1
      }
      break
    }

    if is_prefix {
      fmt.Println("TODO")
      return
    }

    str := string(line)
    value, err1 := strconv.Atoi(str)
    if err1 != nil {
      err = err1
      return
    }

    values = append(values, value)
  }

  return
}

func main() {
  flag.Parse()
  if infile != nil {
    fmt.Println("infile=", *infile, ", outfile=", *outfile, ", algorithm=",
        *algorithm)
  }

  values, err := read_values(*infile)
  if err == nil {
    fmt.Println("Read values", values)
  } else {
    fmt.Println(err)
  }

}
