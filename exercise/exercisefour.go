package main

/*一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。*/
/*一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。*/

import (
	"container/list"
	"fmt"
)

func jumpFloor(n int) int {
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}
	return jumpFloor(n-1) + jumpFloor(n-2)
}

func jumpFloor2(n int) int {
	queue := list.New()
	total := 0
	queue.PushBack(n)
	for queue.Len() > 0 {
		number := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		if number == 1 {
			total = total + 1
		} else if number == 2 {
			total = total + 2
		} else {
			temp1:=number - 1
			temp2:=number - 2
			queue.PushBack(temp1)
			queue.PushBack(temp2)
		}

	}
	return total
}
func jumpFloor3(n int) int {
	if(1==n){
		return 1
	}else if(2==n){
		return 2
	}
	n_1,n_2 ,n_3:= 1,2,0
	for i:=3;i<=n;i++{
		n_3=n_1+n_2
		n_1=n_2
		n_2=n_3
	}
        return n_3

}

func jumpFloorN(n int)int{
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}
	result:=0
	for i:=1;i<n;i++{
		result=result+jumpFloorN(n-i)
	}
	return result
}


func print(result int,no int){
	fmt.Printf("jumpFloor%d: ",no)
	fmt.Println(result)
}
func main() {

	go print(jumpFloor(39),1)
	go print(jumpFloor3(39),3)
	print(jumpFloor2(39),2)

        fmt.Println(jumpFloorN(30))
}
