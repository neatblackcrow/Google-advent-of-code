package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const input = "ckczppom"

func main() {
	for i := 0; i < 1000000000000000000; i += 1 {
		chksum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if chksum[0] == 0 && chksum[1] == 0 && chksum[2] == 0 && chksum[3] == 0 && chksum[4] == 0 {
			fmt.Println(i)
			break
		}
	}
}
