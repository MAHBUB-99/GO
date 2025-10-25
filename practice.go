package main

import (
	"fmt"
	"sync"
)

func main(){
	ff := func (a,b int) {
		c:=a+b
		fmt.Println(c)
	}
	ff(2,3)
    counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

    for i := 0; i < 2; i++ {
		wg.Add(1)
        go func() {
			defer wg.Done()
            for j := 0; j < 10000; j++ {
				mu.Lock()
                counter++
				mu.Unlock()
            }
        }()
    }

	wg.Wait()

    fmt.Println("Counter:", counter)

/*
	var i = 8
	fmt.Printf("%b\n",i)
	fmt.Printf("%b\n",i)
	fmt.Printf("%+d\n",i)
	fmt.Printf("%o\n",i)
	fmt.Printf("%O\n",i)
	fmt.Printf("%x\n",i)
	fmt.Printf("%X\n",i)
	fmt.Printf("%#x\n",i)
	fmt.Printf("%04d\n",i)
	fmt.Println("Hello world")

	// var a bool = true
	// var b int = 5
	// var c float32 = 3.14
	// var d string = "Hi"

	// fmt.Println("Boolean : ",a)
	// fmt.Println("int : ", b)
	// fmt.Println("float32 : ",c)
	// fmt.Println("string : ",d)

	var a  = true
	var b  = 5
	var c  = 3.14
	d  := "Hi"

	fmt.Printf("Boolean : %v and %T \n",a,a)
	fmt.Println("int : ", b)
	fmt.Println("float32 : ",c)
	fmt.Printf("string : %v and %T \n",d,d)

	var test uint8 = 255
	fmt.Println(test)

	var ftest float32 = 3.4e+38
	fmt.Printf("%T and value is : %v \n",ftest,ftest)

	var text1 string = "Hello world"
	var text2 string
	text3 := "New text"
	fmt.Printf("Type: %T and Value is: %v \n",text1,text1)
	fmt.Printf("Type: %T and Value is: %v \n",text2,text2)
	fmt.Printf("Type: %T and Value is: %v \n",text3,text3)

	// var p = 5
	const q = 3.14
	var aa int16 = 10
	const bb  = 5
	fmt.Printf("type: %T and Value: %v \n",q,q)
	fmt.Printf("%v \n",aa+bb)

	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{1,2,3,4,5}
	arr3 := [...]int{5,6,7,8}
	cars := [4]string{"BMW","MAZDA","Volvo","FORD"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(len(arr3))
	cars[3]="LAMBO"
	fmt.Println(cars[3])
	fmt.Println(cars)

	ar1 := [5]int{}
	ar2 := [5]int{1,2,3,4:500}
	ar3 := [5]int{1,2,3,4,5}

	fmt.Println(ar1)
	fmt.Println(ar2)
	fmt.Println(ar3)
	
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	for i,j := 0,len(arr)-1 ; i<j ; i,j = i+1, j-1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
	fmt.Println(arr)

	for i:=0;i<10;i++{
		fmt.Print(arr[i%len(arr)]," ")
	}
		*/
}