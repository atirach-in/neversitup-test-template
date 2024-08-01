package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Permutations(val string) []string {
	if len(val) == 1 {
		return []string{val}
	}

	permutations := make(map[string]struct{})
	for i := 0; i < len(val); i++ {
		str := string(val[i])
		slice := val[:i] + val[i+1:]
		again := Permutations(slice)

		for _, m := range again {
			permutations[str+m] = struct{}{}
		}
	}

	result := make([]string, 0, len(permutations))
	for key := range permutations {
		result = append(result, key)
	}

	return result
}

func FindOdd(arr []int) (string, error) {
	if len(arr) == 0 {
		return "0", nil
	}

	key := make(map[int]int)
	for _, num := range arr {
		key[num]++
	}

	for num, count := range key {
		if count%2 != 0 {
			return strconv.Itoa(num), nil
		}
	}

	return "", errors.New("no odd occurrences found")
}

func CountSmile(arr []string) (int, error) {
	filter := []string{")", "D"}
	count := 0

	for _, element := range arr {
		for _, f := range filter {
			if find := indexOf(element, f); find != -1 {
				count++
				break
			}
		}
	}

	return count, nil
}

func indexOf(str, substr string) int {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func main() {

	result := Permutations("abc")
	fmt.Println("Result Permutations:", result)

	resultOdd, err := FindOdd([]int{1, 1, 2})
	if err != nil {
		fmt.Println("Err FindOdd:", err)
	} else {
		fmt.Println("Result FindOdd:", resultOdd)
	}

	resultCountSmile, err := CountSmile([]string{":)))))", ";(", ";}", ":-DDDDDD"})
	if err != nil {
		fmt.Println("Err CountSmile:", err)
	} else {
		fmt.Println("Result CountSmile:", resultCountSmile)
	}
}
