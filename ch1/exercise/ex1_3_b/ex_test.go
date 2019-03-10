/*
我们用下面的命令运行基准测试。和普通测试不同的是，默认情况下不运行任何基准测试。
我们需要通过-bench命令行标志参数手工指定要运行的基准测试函数。该参数是一个正则表达式，
用于匹配要执行的基准测试函数的名字，默认值是空的。其中“.”模式将可以匹配所有基准测试函数
D:\workspace\go\src\gopl\ch1\exercise\ex1_3_b>go test -bench=.
*/

package main

import (
	"strings"
	"testing"
)

func BenchmarkString2Join(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		result := strings.Join(input, " ")
		if result != "Welcome To China" {
			b.Error("Unexcepted result:" + result)
		}
	}
}

func BenchmarkString2Plus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "China"}
		var s, sep string
		for j := 0; j < len(input); j++ {
			s += sep + input[i]
			sep = " "
		}
		if s != "Welcome To China" {
			b.Error("Unexcepted result:" + s)
		}
	}
}
