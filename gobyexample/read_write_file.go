package main
/*
The io/ioutil example previously worked on the whole file at once.

*/

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main(){
    filename := "rabbits.txt"

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln("Error reading file", filename)
    }

    fmt.Println(string(content))

    outfile := "output.txt"
    err = ioutil.WriteFile(outfile, content, 0777)
    if err != nil {
        log.Fatalln("Error writing file.", err)
    }
}

