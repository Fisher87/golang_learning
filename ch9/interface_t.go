package main

import "fmt"

type Human struct {
    name string
    age int
    phone string
}

func (h Human) SayHi() {
    fmt.Printf("Hi, i am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
    fmt.Println("la la la...", lyrics)
}

type Student struct {
    Human
    school string
    loan float32
}

type Employee struct {
    Human
    company string
    money float32
}

func (e Employee) SayHi() {
    fmt.Printf("Hi, i am %s, i work at %s. call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
    SayHi()
    Sing(lyrics string)
}

func main() {
    mike := Student{Human{"Mike", 25, "12335345"}, "MIT", 0.00}
    paul := Student{Human{"Paul", 26, "12369804"}, "Harvard", 0.00}
    sam := Employee{Human{"sam", 28, "12369804"}, "Golang", 0.00}

    var i Men
    i = mike
    i.SayHi()
    i.Sing("rain")

    i=sam
    i.SayHi()
    i.Sing("wild")

    x := make([]Men, 3)
    x[0], x[1], x[2] = paul, sam, mike

    for _, value := range x {
        value.SayHi()
    }

}
