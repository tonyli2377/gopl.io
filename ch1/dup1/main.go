// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	//从程序的标准输入中读取内容
	input := bufio.NewScanner(os.Stdin)

	//每次调用 input.Scan() ，即读入下一行，并移除行末的换行符；读取的内容可以调用 input.Text() 得到。
	//Scan()在读到一行时返回 true ，不再有输入时返回 false 。
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()

	for line, n := range counts {
		//println(n)
		if n > 1 {
			//fmt.Printf 函数对一些表达式产生格式化输出。该函数的首个参数是个格式字符串,
			//指定后续参数被如何格式化。各个参数的格式取决于“转换字符”（conversion character），形式为百分号后跟一个字母。
			// %d 表示以十进制形式打印一个整型操作数，而 %s 则表示把字符串型操作数的值展开。
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
