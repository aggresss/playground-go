package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []int64{-65, 1, 2, 127, 128, 255, 256} {
		n := binary.PutVarint(buf, x)
		fmt.Print(x, " var length: ", n, ", hex: 0x")
		fmt.Printf("%x\n", buf[:n])
	}

}
