package main

import "fmt"

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func ExtendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := ExtendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

func main() {
	gcd := GCD(35, 15)
	fmt.Printf("The GCD of 35 and 15 is %d\n", gcd)
	gcd, x, y := ExtendedGCD(35, 15)
	fmt.Printf("The GCD of 35 and 15 is %d, x = %d, y = %d\n", gcd, x, y)
}
