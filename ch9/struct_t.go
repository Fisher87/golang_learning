package main

import "fmt"

type Count int //创建自定义类型 Count

func (count *Count) InCrement() { *count++ }
func (count *Count) Decrement() { *count-- }
func (count Count) IsZero() bool { return count==0 }

type Part struct {
    stat string
    Count
}

func (part Part) IsZero() bool {     //覆盖了匿名字段Count的IsZero()方法
    return part.Count.IsZero() && part.stat == ""
}

func (part Part) String() string {
    return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
    var i Count = -1
    fmt.Printf("Start \"Count\" test:\n Origin value of count: %d\n", i)
    i.InCrement()
    fmt.Printf("value of count after increment: %d\n", i)
    fmt.Printf("count is zero t/f? : %t\n\n", i.IsZero())
    fmt.Println("start: Part test")
    part := Part{"232", 0}
    fmt.Printf("Part: %v\n", part)
    fmt.Printf("Part is zero t/f? :%t\n", part.IsZero())
    fmt.Printf("Count in part is zero t/f?: %t\n", part.Count.IsZero())
}
