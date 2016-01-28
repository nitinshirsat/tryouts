package tryouts

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

type map_str_int map[string]int

func Dup ()  {
    dup_count := make(map_str_int)

    files := os.Args[2:]
    if len(files) < 1 {return }
    for _,filename := range(files){

        fmt.Printf("Processing file %s \n",filename)
        file_content, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr,"Error occured while reading file : %v \n error message:%v ",err,)
        }
        content_lines :=  strings.Split(string(file_content),"\n")
        for _, line := range(content_lines){
            dup_count[line]++
        }

        fmt.Printf("Duplicate lines are :\n")
        // print all duplicate lines
        for key, val:= range (dup_count){
            if val > 1 {
                fmt.Printf("%d : %s \n",val,key)
            }
        }
    }
}
