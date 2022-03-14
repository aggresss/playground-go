package main

import "fmt"

func main() {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(RemoveDuplicationElement(arr))
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

func RemoveDuplicationElement(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		if _, ok := set[v]; ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
