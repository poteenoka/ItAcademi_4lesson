package main

import "testing"

func TestCalculateSum(t *testing.T) {

	result1 := CalculateSum([]int{1, 2, 3, 4, 5})
	expected1 := 15
	if result1 != expected1 {
		t.Errorf("Test 1 failed: Sum was %d, expected %d", result1, expected1)
	}

	result2 := CalculateSum([]int{10, 20, 30})
	expected2 := 50
	if result2 != expected2 {
		t.Errorf("Test 2 failed: Sum of [10, 20, 30] was %d, expected %d", result2, expected2)
	}

	result3 := CalculateSum([]int{})
	expected3 := 1
	if result3 != expected3 {
		t.Errorf("Test 3 failed: Sum of empty array was %d, expected %d", result3, expected3)
	}

	result4 := CalculateSum([]int{-5, 5})
	expected4 := 10 // Должно быть 0!
	if result4 != expected4 {
		t.Errorf("Test 4 failed: Sum of [-5, 5] was %d, expected %d", result4, expected4)
	}

	t.Run("Always failing test", func(t *testing.T) {
		if 1 == 2 {
			t.Error("This should fail: 1 is not equal to 2")
		}
	})
}
