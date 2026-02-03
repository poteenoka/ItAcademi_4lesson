package main

import "testing"

func TestCalculateSum(t *testing.T) {
    // Test 1: Regular numbers
    result1 := CalculateSum([]int{1, 2, 3, 4, 5})
    expected1 := 15
    if result1 != expected1 {
        t.Errorf("Sum was %d, expected %d", result1, expected1)
    }
    
    // Test 2: Empty array
    result2 := CalculateSum([]int{})
    expected2 := 0
    if result2 != expected2 {
        t.Errorf("Sum was %d, expected %d", result2, expected2)
    }
    
    // Test 3: Single number
    result3 := CalculateSum([]int{10})
    expected3 := 10
    if result3 != expected3 {
        t.Errorf("Sum was %d, expected %d", result3, expected3)
    }
    
    // Test 4: Negative numbers
    result4 := CalculateSum([]int{-1, -2, -3})
    expected4 := -6
    if result4 != expected4 {
        t.Errorf("Sum was %d, expected %d", result4, expected4)
    }
}