package main

import "fmt"

func main() {
    fmt.Println("The extra-numbers program shows you how to add, subtract")
    fmt.Println("multiply and divide numbers and how to use brackets.")
    fmt.Println("Sums without brackets.")
    fmt.Print("3 + 2 * 5 = ")
    fmt.Println(3 + 2 * 5)
    fmt.Print("4 * 4 + 2 = ")
    fmt.Println(4 * 4 + 2)

    fmt.Print("25 - 20 / 4 = ")
    fmt.Println(25 - 20/4)

    fmt.Print("30 / 2 + 5 = ")
    fmt.Println(30 / 2 + 5)

    fmt.Println("Sums with brackets.")
    fmt.Print("3 + (2 * 5) = ")
    fmt.Println(3 + (2 * 5))

    fmt.Print("(3 + 2) * 5 = ")
    fmt.Println((3 + 2) * 5)

    fmt.Print("4 * (4 + 2) = ")
    fmt.Println(4 * (4 + 2))

    fmt.Print("(28 - 20) / 4 = ")
    fmt.Println((28 - 20) / 4)

    fmt.Print("30 / (2 + 5 + 3) = ")
    fmt.Println(30 / (2 + 5 + 3))
}