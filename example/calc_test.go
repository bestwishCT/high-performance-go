package main

import (
	"fmt"
	"testing"
)

//func TestAdd(t *testing.T) {
//	if ans := Add(1, 2); ans != 3 {
//		t.Errorf("1 + 2 expected be 3, but %d got", ans)
//	}
//
//	if ans := Add(-10, -20); ans != -30 {
//		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
//	}
//}

//func TestMul(t *testing.T) {
//	t.Run("pos", func(t *testing.T) {
//		if Mul(2, 3) != 6 {
//			t.Fatal("fail")
//		}
//
//	})
//	t.Run("neg", func(t *testing.T) {
//		if Mul(2, -3) != -6 {
//			t.Fatal("fail")
//		}
//	})
//}
// 3 子测试(Subtests)
// 对于多个子测试的场景，更推荐如下的写法(table-driven tests)：
//func TestMul(t *testing.T) {
//	cases := []struct {
//		Name           string
//		A, B, Expected int
//	}{
//		{"pos", 2, 3, 6},
//		{"neg", 2, -3, -6},
//		{"zero", 2, 0, 0},
//	}
//
//	for _, c := range cases {
//		t.Run(c.Name, func(t *testing.T) {
//			if ans := Mul(c.A, c.B); ans != c.Expected {
//				t.Fatalf("%d * %d expected %d, but %d got",
//					c.A, c.B, c.Expected, ans)
//			}
//		})
//	}
//}

//4 帮助函数(helpers)
/*
*
关于 helper 函数的 2 个建议：
不要返回错误， 帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性。
调用 t.Helper() 让报错信息更准确，有助于定位。
*/
//type calcCase struct{ A, B, Expected int }
//
//func createMulTestCase(t *testing.T, c *calcCase) {
//	t.Helper()
//	if ans := Mul(c.A, c.B); ans != c.Expected {
//		t.Fatalf("%d * %d expected %d, but %d got",
//			c.A, c.B, c.Expected, ans)
//	}
//
//}
//
//func TestMul(t *testing.T) {
//	createMulTestCase(t, &calcCase{2, 3, 6})
//	createMulTestCase(t, &calcCase{2, -3, -6})
//	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
//}

// 5 setup 和 teardown
//
//	func setup() {
//		fmt.Println("Before all tests")
//	}
//
//	func teardown() {
//		fmt.Println("After all tests")
//	}
//
//	func Test1(t *testing.T) {
//		fmt.Println("I'm test1")
//	}
//
//	func Test2(t *testing.T) {
//		fmt.Println("I'm test2")
//	}
//
//	func TestMain(m *testing.M) {
//		setup()
//		code := m.Run()
//		teardown()
//		os.Exit(code)
//	}
//
// 6 网络测试(Network)
//
//	func TestArray(t *testing.T) {
//		a := [...]int{1, 2, 3}
//		b := a
//		a[0] = 100
//		fmt.Println(a)
//		fmt.Println(b)
//	}
func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

//	func TestSliceLenAndCap(t *testing.T) {
//		nums := []int{1}
//		printLenCap(nums) // len: 1, cap: 1 [1]
//		nums = append(nums, 2)
//		printLenCap(nums) // len: 2, cap: 2 [1 2]
//		nums = append(nums, 3)
//		printLenCap(nums) // len: 3, cap: 4 [1 2 3]
//		nums = append(nums, 3)
//		printLenCap(nums) // len: 4, cap: 4 [1 2 3 3]
//		nums = append(nums, 4)
//		printLenCap(nums)
//	}
func TestSlice(t *testing.T) {
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]
}

/*
*
在已有切片的基础上进行切片，不会创建新的底层数组。因为原来的底层数组没有发生变化，内存会一直占用，直到没有变量引用该数组。
因此很可能出现这么一种情况，原切片由大量的元素构成，但是我们在原切片的基础上切片，虽然只使用了很小一段，但底层数组在内存中仍然占据了大量空间，得不到释放。
比较推荐的做法，使用 copy 替代 re-slice。
*/
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}
