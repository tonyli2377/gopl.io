package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	LowEf()
	HighEf()
}

func LowEf() {
	arrStr := [...]string{"a", "b", "c", "d", "e", "f"}
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		tempStr := ""
		for _, str := range arrStr {
			tempStr += str + ","
		}
		//fmt.Println(tempStr)
	}
	fmt.Printf("%.2fs LowEf elapsed\n", time.Since(start).Seconds())
}

func HighEf() {
	arrStr := [...]string{"a", "b", "c", "d", "e", "f"}
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		//fmt.Println(strings.Join(arrStr[:], ","))
		strings.Join(arrStr[:], ",")
	}
	fmt.Printf("%.2fs HighEf elapsed\n", time.Since(start).Seconds())
}
