package main

import "fmt"

func INSERTIONSORT(A []int, m int, p int) {
	i,j:=0,0
	for j=m+1;j<=p;j++{
		item:=A[j]
		i=j-1
		for item<A[i]{
			A[i+1]=A[i]
			i=i-1
			if i<m{
				break
			}
		}
		A[i+1]=item
	}
}
/**
A 数组
m 起始下标
p 元素个数
 */
func PARTITION(A []int, m int, p int) int {
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
			INTERCHANGE(&A[i], &A[p])
		}
	}
	A[m] = A[p]
	A[p] = v
	return p
}

func INTERCHANGE(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
/**
A 数组
m 起始下标
p 结尾下标
k 第k小
r 分组长度
 */
func SEL(A []int, m int, p int, k int, r int) int {

	n, i, j := 0, 0, 0
	if p-m+1 <= r {
		INSERTIONSORT(A, m, p)
		return m + k - 1
	}
	for {
		n = p - m + 1
		for i = 1; i <= n/r; i++ {
			INSERTIONSORT(A, m+(i-1)*r, m+i*r-1)
			INTERCHANGE(&A[m+i-1], &A[m+(i-1)*r+r/2])
		}
                k1:=(n/r)/2
		if k1==0{
			k1=1
		}
		j = SEL(A, m, m+n/r-1, k1, r)
		INTERCHANGE(&A[m], &A[j])
		j = p + 1
		j=PARTITION(A, m, j)
		if j-m+1 == k {
			return j
		} else if j-m+1 > k {
			p = j - 1
		} else {
			k = k - (j - m + 1)
			m = j + 1
		}
	}
}

func main() {
	A := []int{8, 6, 9, 3, 5, 2, 7,11,2,15,1,0,24}
	i:=SEL(A,0,12,8,3)
	fmt.Println(i)
	fmt.Println(A)
	fmt.Println(A[i])

}
