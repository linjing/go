package main

import "fmt"

func my_func(args ... int) {
    for i, arg := range args {
        fmt.Println("%d, %d", i, arg)
    }
}
func my_func2(args [] int) {
    for i, arg := range args {
        fmt.Println("%d, %d", i, arg)
    }
}

func my_print(args ...interface{}) {
    for _, arg := range args {
      switch arg.(type) {
      case int: fmt.Println(arg, "is int")
    case string: fmt.Println(arg, "is string")
        case int64: fmt.Println(arg, "is i64")
    default: fmt.Println(arg, "unknown")
        }
    }
}

func main() {
    fmt.Println("bb");
    my_func(1,2,3,3,2,1)
    fmt.Println("bb");
    my_func2([] int {1,2,3,3,2,1})
    my_print(3,"11")

}
