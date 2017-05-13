package main

import "fmt"

func main() {
    x := uint16(6500)
    y := int16(x)
    fmt.Printf("type and value of x is: %T and %d \n", x, x)
    fmt.Printf("type and value of y is: %T and %d \n", y, y)

    var i interface{} = 99
    /* var s interface{} = []string{"left", "right"} */

    /* i := i.(type) */
    switch i := i.(type) {
    case int:
        fmt.Printf("type and value of j is: %T and %d \n", i, i)
    }
}
