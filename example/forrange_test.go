package main

import (
	"fmt"
	"testing"
	"time"
)

/*
*
变量 words 在循环开始前，仅会计算一次，如果在循环中修改切片的长度不会改变本次循环的次数。
迭代过程中，每次迭代的下标和值被赋值给变量 i 和 s，第二个参数 s 是可选的。
针对 nil 切片，迭代次数为 0
*/
func TestForRange(t *testing.T) {
	words := []string{"Go", "语言", "高性能", "编程"}
	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}
}

func TestForRanageChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "Go"
		ch <- "语言"
		ch <- "高性能"
		ch <- "编程"
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

func TestRangeCopy(t *testing.T) {
	persons := []struct{ no int }{{no: 1}, {no: 2}, {no: 3}}
	for _, s := range persons {
		s.no += 10
	}
	for i := 0; i < len(persons); i++ {
		persons[i].no += 100
	}
	fmt.Println(persons) // [{101} {102} {103}]
}

// 1: 切片扩展陷阱
/*
解析：
append 操作： Go 中的 append 函数用于在切片末尾添加元素。如果切片的容量 (cap) 足够容纳新的元素，append 会直接扩展切片并更新其长度；如果容量不足，Go 会分配一个更大的底层数组，并将元素复制过去。

容量的增长： Go 的切片扩展策略通常是容量会按以下规则增长：

如果当前容量为 n，append 会将切片容量扩展为 2n，这种增长方式能在大多数情况下提高效率。

具体的增长方式依赖于 Go 的内部实现，通常是根据当前容量的大小来选择合适的增长量。
*/
func TestSliceAppend(t *testing.T) {
	var a []int
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))
	}
}

// 2: 多线程修改同一切片
func TestConcurrency(t *testing.T) {
	slice := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		go func() {
			slice[0]++
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(slice)
}

// 3: map 并发读写
func TestMapConcurrency(t *testing.T) {
	m := make(map[string]int)

	// 5个协程并发写入
	for i := 0; i < 5; i++ {
		go func(i int) {
			m[fmt.Sprintf("key%d", i)] = i
		}(i)
	}

	// 5个协程并发读取
	for i := 0; i < 5; i++ {
		go func(i int) {
			_ = m[fmt.Sprintf("key%d", i)]
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(m)
}

// 4 指针与切片
/**
解析：
切片的赋值： 在 Go 中，切片的赋值并不会创建新的切片，而是将 底层数组的引用 赋给新切片。所以 a 和 b 共享同一个底层数组。

修改效果： 当你修改 b 中的值时，因为 a 和 b 指向同一个底层数组，所以 a 中的值也会受到影响。
*/
func TestPointerSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
