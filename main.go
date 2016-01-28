package main

import (
   "os"
    "fmt"
    "tryouts/try"
)

func main(){

    for _, arg := range(os.Args[1:]){
        switch {
        case arg == "dup":
            tryouts.Dup();

        case arg =="echo":
            tryouts.Echo()

        case arg == "concurrancy":
            tryouts.Concurrancy()

        default:
         fmt.Println("please confirm parameter, it does not match")
        }
    }
}
