//Ftoc prints two 
package main

import "fmt"

/* func main() { */
/*     const freezingF, boilingF = 32.0, 212.0 */
/*     var c1, c2 float64 */
/*     c1 = fToc(freezingF) */
/*     c2 = fToc(boilingF) */
/*     fmt.Printf("%g°F = %g°C\n", freezingF, c1) */
/*     fmt.Printf("%g°F = %g°C\n", boilingF, c2) */

/* } */

/* func fToc(f float64) (c float64) { */
/*     c = (f - 32) * 5 / 9 */
/*     return c */
/* } */

func main() {
    v := 1
    incr(&v)
    fmt.Println(v)

}

func incr(v *int) int {
    *v = *v+1
    return *v
}
