package recursion

import "errors"

// 1: 1; 2: 1 1, 2; f(3)= f(2)+f(1)
// 递推公式：f(1)=1, f(2)=2, f(n)=f(n-1)+f(n-2)
// 退出条件：n=1 || n=2 return 1 || 2

var m map[int]int

func calcCore(n int) (num int) {
	if n <= 2 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	m[n-1] = calcCore(n - 1)
	m[n-2] = calcCore(n - 2)
	return m[n-1] + m[n-2]
}

func CalWaysNum(n int) (num int, err error) {
	if n < 1 {
		return 0, nil
	}
	if n > 2048 {
		return 0, errors.New("input n < 2048!")
	}
	m = make(map[int]int, n)
	return calcCore(n), nil
}
