package main

import (
	"fmt"
)

func patition(A []int, m int, p int) int {
	v, i := A[m], m
	for i < p {
		i++
		p--
		for {
			if A[i] >= v {
				break
			} else {
				i++
			}

		}
		for {
			if A[p] <= v {
				break
			} else {
				p--
			}
		}
		if i < p {
			tmp := A[i]
			A[i] = A[p]
			A[p] = tmp
		}
	}
	A[m] = A[p]
	A[p] = v
	return p
}

func quicksort(A []int, p int, q int) {
	var stack [20]int
	j,top := 0,0
	for {
		for p < q {
			j = q + 1
			j = patition(A, p, j)
			if j - p < q - j {
				stack[top+1]=j+1
				stack[top+2]=q
				q = j - 1
			} else {
				stack[top+1]=p
				stack[top+2]=j - 1
				p = j + 1
			}
			top=top+2
		}
		if top==0{
			break
		}
                q=stack[top]
		p=stack[top-1]
		top=top-2
	}

}

func main() {
	B := []int{10, 9, 8, 7, 6, 5, 4,3,2,1}
	patition(B ,0,12)
	fmt.Println(B)
	/*A := []int{8, 6, 9, 3, 5, 2, 7,11,2,15,1,0,24}
	quicksort(A,0,12)
	fmt.Println(A)*/

}
