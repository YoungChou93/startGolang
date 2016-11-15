package main
/* 计算Fibonacci第n个数 */

import "fmt"

type LongNumber struct {
	number []int
}

func newLongNumber(number []int) *LongNumber {
	var result = new(LongNumber)
	result.number = number
	return result
}

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func (n1 *LongNumber) add(n2 *LongNumber) {
	for index, _ := range n1.number {
		if n1.number[index] != 0 || n2.number[index] != 0 {
			n1.number[index] = n1.number[index] + n2.number[index]
			if n1.number[index] >= 100000 {
				temp1 := n1.number[index] % 100000
				temp2 := n1.number[index] / 100000
				n1.number[index] = temp1
				n1.number[index+1] = n1.number[index+1] + temp2
			}
		} else {
			break
		}
	}
}
func (ln *LongNumber) print() {
	for i := len(ln.number) - 1; i >= 0; i-- {
		if 0 != ln.number[i] {
			fmt.Print(ln.number[i])
			for j := i-1; j >= 0; j-- {
				fmt.Printf("%0[1]*.[2]*[3]d" ,5,5,ln.number[j])
			}
			break
		}
	}
	fmt.Println()

}

func FibonacciLong(n int) *LongNumber {

	if n == 1 || n == 2 {
		number := []int{50: 0}
		number[0] = 1
		return newLongNumber(number)
	}

	result1 := FibonacciLong(n - 1)
	result2 := FibonacciLong(n - 2)
	result1.add(result2)

	return result1

}
func Fibadd(n1 []int, n2 []int) []int {
	 result:= []int{100:0}
	length := len(n1)
	if length < len(n2) {
		length = len(n2)
	}
	for i := 0; i < length; i++ {
		if n1[i] == 0 && n2[i] == 0 {
			break
		}
		result[i] = result[i]+n1[i] + n2[i]
		if result[i] >= 100000 {
			temp1 := result[i] / 100000
			temp2 := result[i] % 100000
			result[i] = temp2
			result[i+1] = result[i+1] + temp1
		}
	}

	return result

}
func FibonacciLong2(n int) []int {
	var fib [10][]int
	for i := 0; i < 10; i++ {
		fib[i] = []int{100: 0}
	}
	if n == 0 {
		return fib[0]
	}

	fib[0][0] = 1
	fib[1][0] = 1

	if n == 1 || n == 2 {
		return fib[0]
	}

	n = n - 1
	for i := 2; i <= n; i++ {
		fib[i % 10] = Fibadd(fib[(i-1) % 10],fib[(i-2) %10])
	}

	return fib[n%10]
}
func main() {
	fmt.Println(Fibonacci(30))
	FibonacciLong(30).print()
	newLongNumber(FibonacciLong2(30)).print()

	newLongNumber(FibonacciLong2(2000)).print()

}
