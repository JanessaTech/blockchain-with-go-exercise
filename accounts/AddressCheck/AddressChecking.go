package addresscheck

import (
	"fmt"
	"regexp"
)

func CheckPublicAddress() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	addr1 := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	addr2 := "0xZYXb5d4c32345ced77393b3530b1eed0f346429d"

	fmt.Printf("%s is valid: %v\n", addr1, re.MatchString(addr1))
	//0x323b5d4c32345ced77393b3530b1eed0f346429d is valid: true
	fmt.Printf("%s is valid: %v\n", addr2, re.MatchString(addr2))
	//0xZYXb5d4c32345ced77393b3530b1eed0f346429d is valid: false
}
