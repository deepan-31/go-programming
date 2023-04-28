package main

import "fmt"

func main(){
/*
	fmt.Println("hello world")

	var age int = 40

	var favNum float64 = 1.414

	randNum := 1;

	fmt.Println(age,favNum,randNum)
	const pi float64 = 3.14156

	//var myName string ="Derek Banas"

	//fmt.Printf(myName)

	fmt.Printf("%.3f \n",pi)

	fmt.Printf("%T \n",pi)

	fmt.Printf("%e \n",pi)
	
	fmt.Println("true && false=",true && false)
	fmt.Println("true || false=",true || false)
	fmt.Println("not true=",!true)
	
	i :=10
	for i>= 0 {
		fmt.Println(i)

		i--
	}
	
	for j:=0;j<5;j++{
		fmt.Println(j)
	}
	
	yourAge := 18
	if yourAge >=16{
		fmt.Println("you can drive")
	}else{
		fmt.Printf("You can drive \n")
	}
	
	yourAge := 5
	switch yourAge {
	case 16 : fmt.Println("Go drive")
	case 18 : fmt.Println("Go Vote")
	default : fmt.Println("Go have fun")
	}
	
	//var favNums2 = [...]int{1,2,3,4,5}

	favNums3 := [5]float64{1,2,3,4,5}

	for _, value := range favNums3{
		fmt.Println(value)
	}
	
	numSlice := []int{5,4,3,2,1}
	numSlice2 := numSlice[3:5]

	fmt.Println("numSlice2 = ",numSlice2)
	

	fmt.Println("numSlice[:2]",numSlice[:2])
	

	presAge := make(map[string] int)

	presAge["theodreRosevolt"] = 42
	presAge["JohnF.Kennedy"] = 43
	fmt.Println(len(presAge))

	delete(presAge,"theodreRosevolt")
	fmt.Println(len(presAge))

	
	listNums := []float64{1,2,3,4,5}

	sum := addThemUp(listNums)

	fmt.Println(sum)
	
	num1,num2 := next2Values(5)
	fmt.Println(num1,num2)

	*/
	fmt.Println(subtractThem (1,2,3,4,5))
}
/*
func addThemUp(numbers []float64) float64{
	sum := 0.0
	for _, val := range numbers{
		sum += val 
	}
	return sum
}

func next2Values(number int) (int,int){
	return number+1,number+2
}
*/
func subtractThem(args ...int) int{
	finalValue:= 0 

	for _,value := range args{
		finalValue -= value
	}

	return finalValue
}
