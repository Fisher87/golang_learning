package main

import "fmt"

func classchecker(items ...interface{}) { //build a function which can revise all kind params.
    for i, x := range items {
        switch x := x.(type) {
        case bool:
            fmt.Printf("param #%d is a bool, value: %t\n", i, x)
        case int:
            fmt.Printf("param #%d is a int, value: %d\n", i, x)
        }
    }

}

func main() {
    classchecker(5, true)
}
