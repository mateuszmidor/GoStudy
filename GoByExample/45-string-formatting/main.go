package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)                      //{1 2}
	fmt.Printf("%+v\n", p)                     // {x:1 y:2}
	fmt.Printf("%#v\n", p)                     // main.point{x:1, y:2}
	fmt.Printf("%T\n", p)                      // main.point
	fmt.Printf("%t\n", true)                   // true
	fmt.Printf("%d\n", 123)                    // 123
	fmt.Printf("%b\n", 14)                     // 1110
	fmt.Printf("%c\n", 33)                     // !
	fmt.Printf("%X\n", 456)                    // 1C8
	fmt.Printf("%.2f\n", 36.6)                 // 36.60
	fmt.Printf("%e\n", 1234000.0)              // 1.234000e+06
	fmt.Printf("%s\n", `"string"`)             // "string"
	fmt.Printf("%q\n", "string")               // "string"
	fmt.Printf("% x\n", "123 XYZ")             // 31 32 33 20 58 59 5a
	fmt.Printf("%p\n", &p)                     // 0xc0000140f0
	fmt.Printf("|%6d|%6d|\n", 12, 345)         // |    12|   345|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   // |  1.20|  3.45|
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |
	fmt.Printf("|%6s|%6s|\n", "foo", "b")      // |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")    // |foo   |b     |
	fmt.Printf("|%*s5 spaces|\n", 5, "")       // |     5 spaces|
}
