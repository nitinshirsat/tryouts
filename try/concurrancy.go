package tryouts

import (
    "fmt"
    "time"
)

func Concurrancy(){

    var mylist []int

    for i:= 0; i< 30; i++ {

        mylist = append(mylist, i)
    }
    num_type := make (chan string, 4)

    fmt.Println("welcome to concurrancy example")

    for _,i := range (mylist){

        switch {
        case i== 0 :
            fmt.Println("0 Wow not odd not even")
            break

        case i%2 ==0 :
            time.Sleep(time.Duration(i*10000000))
            go even_odd(i, num_type)
            num_type <- "Even number"
            break

        default:

            go even_odd(i, num_type)
            num_type <- "odd number"
            break;
        }

    }
}

// num_type chan string ,

func even_odd(i int, num_type chan string ){

    fmt.Printf("%d  is %s\n", i, <-num_type)
    time.Sleep(time.Duration(i*10000))
}

