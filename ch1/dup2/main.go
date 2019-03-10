// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			//os.Open 函数返回两个值。第一个值是被打开的文件( *os.File ），其后被 Scanner 读取。
			//第二个值是内置 error 类型的值。
			f, err := os.Open(arg)
			if err != nil {
				//向标准错误流打印一条信息
				//%v任意类型默认格式,
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				//跳到 for 循环的下个迭代开始执行
				continue
			}
			//countLines 函数在其声明前被调用。函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用。
			//map 作为参数传递给某函数时，该函数接收这个引用的一份拷贝。
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			//fmt.Println()
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
