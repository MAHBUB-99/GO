package main
import ("fmt")

func main(){
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
}