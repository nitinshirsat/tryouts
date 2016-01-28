package tryouts

import (

    "fmt"
    "os"
    "bufio"

)

type int_map map[string]int

func Echo() {

    count := make(int_map)

    files := os.Args[2:]
    var fp *os.File
    var err error
    fmt.Print(files)
    if len(files) > 0 {
        //get the file first
        for _,filename := range(files){

            fp, err = os.Open(filename)

            if err != nil {
                fmt.Fprintf(os.Stderr,"%v",err)
                return
            }

        }
    } else {
        fp = os.Stdin
    }
    countLines(fp, count)

    var s string;
    s = "Count \t String "
    for str, count := range count{
        if count > 1 {
            s += " "+ fmt.Sprintf(" %v: \t %s \n",count, str)
        }

    }

    fmt.Println(s)
}

func countLines(filep *os.File, count int_map){

    input := bufio.NewScanner(filep)
    for input.Scan(){
        count[input.Text()]++
    }

}

