package main

func main() {
	a := 5
	iPanic(a, &a)
}

func iPanic(i int, j *int) {
	if i > 0 {
		iPanic(i-1, j)
	}
	panic("panic here")
}
