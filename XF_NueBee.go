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

	//GenRsaKey(4096)

	fmt.Println("---------------RSA-----------------")
	bmsg, err := RsaEncrypt([]byte(nnn))
	if err == nil {
		fmt.Println(bmsg)
	} else {
		fmt.Println("error 1:", err.Error())
	}
	omsg, err := RsaDecrypt(bmsg)
	if err == nil {
		fmt.Println(string(omsg))
		fmt.Println(len(omsg))
	} else {
		fmt.Println("error 2:", err.Error())
	}
}