package main

import "fmt"

// // 基本用法1: 资源清理
// func f1() {
// 	defer fmt.Println("close 1")
// 	defer fmt.Println("close 2")

// 	fmt.Println("3")
// }

// // 基本用法2：异常恢复
// func f2() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("paniced: %+v \n", err)
// 		}
// 	}()
// 	panic("test")
// }

// runtime.deferprocStack(SB)
// 仍然是在栈上，返回前修改 res 返回 res
func f3() (res int) {
	defer func() {
		res++
	}()
	return 0
}

// runtime.deferprocStack(SB)
// 参数复制, 不会影响返回值
// func f4() (res int) {
// 	defer func(res int) {
// 		res++
// 	}(res)
// 	return 0
// }

// 如何判断 defer 在堆上还是在栈上呢？
// https://github.com/golang/go/blob/master/src/cmd/compile/internal/gc/escape.go#L743
// 主要是在逃逸分析的时候，发现 e.loopDepth == 1  并且不是 open-coded defer 就会分配到栈上。
// 这也是为什么 go 1.13 之后 defer 性能提升的原因，所以切记不要在循环中使用 defer 不然优化也享受不到

// runtime.deferproc(SB) 堆分配
func f6() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("f6: %d\n", i)
		}()
	}
}
func main() {
	// 栈
	r := f3()
	fmt.Println(r)

	// 栈
	// fmt.Println(f4())
}
