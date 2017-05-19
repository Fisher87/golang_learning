package main

import "fmt"

type Dog struct {
    Name string
    Color string
}

func (d Dog) Call() {
    fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

func main() {
    Spot := Dog{Name:"spot", Color:"brown"}
    var Rover Dog
    Rover.Color = "light"
    Rover.Name = "Rover"

    Spot.Call()
    Rover.Call()
}
