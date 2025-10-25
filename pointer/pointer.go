package main
import "fmt"


func updateValue(num *int) {
	*num = *num+10
}

func updateSlice(nums []int) {
	nums[0]=100
}

func main () {
	nums:=[]int{1,2,3}
	updateSlice(nums)
	fmt.Println(nums)

	x:=5
	updateValue(&x)
	fmt.Println(x)
}