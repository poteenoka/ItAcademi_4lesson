package main

import "fmt"

// CalculateSum calculates sum of numbers in array
func CalculateSum(numbers []int) int {
    sum := 0
    for i := 0; i < len(numbers); i++ {
        sum += numbers[i]
    }
    return sum
}

func main() {
    // Example array
    numbers := []int{1, 2, 3, 4, 5}
    
    result := CalculateSum(numbers)
    fmt.Println("Numbers:", numbers)
    fmt.Println("Sum:", result)
}