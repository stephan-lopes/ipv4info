package main

import (
	"fmt"
)

func main() {
	sliceIP, err := readFile("../ip.txt")
	if err != nil {
		panic(err)
	}

	sliceIP = createLink(sliceIP)

	for _, ip := range sliceIP {
		country, err := checkIP(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(country)
	}
}
