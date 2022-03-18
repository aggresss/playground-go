package main

import "fmt"

func main() {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(RemoveRepeatedElement(arr))
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// DeduplicateStringSlice deduplicate string slice
func DeduplicateStringSlice(input []string) (output []string) {
	set := make(map[string]struct{}, len(input))
	j := 0
	for _, v := range input {
		if _, ok := set[v]; ok {
			continue
		}
		set[v] = struct{}{}
		input[j] = v
		j++
	}
	output = input[:j]

	return output
}
