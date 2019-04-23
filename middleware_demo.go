// 中间件框架思想实践
package main 

import (
	"fmt"
)

type calcHandler func(x int) int

type middleware func(calcHandler) calcHandler

func main() {
	allInOne := chain(mw1, mw2, mw3)
	allInOne(calc)(3)
}

func calc(d int) int {
	fmt.Println("calc handle complete, result =", d + 1)
	return d + 1
}

func mw1(next calcHandler) calcHandler {
	return func(x int) int {
		fmt.Println("mw1 work")
		return next(x)
	}
}

func mw2(next calcHandler) calcHandler {
	return func(x int) int {
		fmt.Println("mw2 work")
		return next(x)
	}
}

func mw3(next calcHandler) calcHandler {
	return func(x int) int {
		fmt.Println("mw3 work")
		return next(x)
	}
}

func chain(outer middleware, others ...middleware) middleware {
	return func(next calcHandler) calcHandler {
		for i := len(others) - 1; i >= 0; i -- {
			next = others[i](next)
		}
		return outer(next)
	}
}