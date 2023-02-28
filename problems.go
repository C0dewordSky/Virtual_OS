// Prime numbers in a given range

package main

import (
	"fmt"
	"math"
)


func printPrime(num1, num2 int){

	if num1< 2 || num2<2{
	fmt.Println("Enter an number greater than 2")
	return
	}

	for num1<=num2{
	isPrime := true
	

	for i:=2;i<=int(math.Sqrt(float64(num1)));i++{
	if num1 % i == 0{
		isPrime=false;
		break
	}
}
     if isPrime{
		fmt.Printf("%d ", num1)
	}
	num1++
}
    fmt.Println()
     
	}


func main(){
	printPrime(4,16)
}


// Bar chart 

func barChart(){

	
}