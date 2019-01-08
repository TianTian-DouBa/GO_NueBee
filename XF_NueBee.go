package main

import (
	"fmt"
	//"net"
	. "./xf_utility"
)

func main() {
	lll := MakeATxt()
	fmt.Println("coated:",lll)

	nnn, _ := OpenATxt(lll)
	fmt.Println("original:", nnn)

	GenRsaKey(1024)
}