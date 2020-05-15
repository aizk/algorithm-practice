package main

import "log"

func main() {
	log.Println(resolve('a'))
	log.Println(resolve(0xF))
	log.Println(resolve2(0xFFF))
}

// 循环右移 1 - 7 位
// 与 1 作 & 运算判断第 x 位是不是 1
func resolve_(x uint8) int {
	count := 0
	for i := 0; i < 8; i++ {
		if (x >> uint8(i)) & 1 > 0 {
			count++
		}
	}
	return count
}

// 修改版
// 之前的写法太荣誉了，没想到可以写得如此简单...
// & 运算完之后的数字可以直接拿过来使用...
func resolve(x int) (count int) {
	for x > 0 {
		count += x & 0x01
		x >>= 1
	}
	return
}

// --- 书上的解法

// 除 2 法
func resolve2(x int) (count int) {
	count = 0
	for x > 0 {
		if x%2 == 1 {
			count++
		}
		x = x / 2
	}
	return
}