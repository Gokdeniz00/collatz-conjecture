package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
)

func main(){
	var number int
	fmt.Print("Enter a number for Collatz Conjecture:")
	fmt.Scan(&number)
	values := arraylist.New()
	Collatz(number, values,true)
	fmt.Print(values)

}
func Collatz(number int,values *arraylist.List,first bool) {
	if first{
		values.Add(number)
	}
	if number % 2 == 0 {
		number = number/2
		values.Add(number)
	}else if number % 2 == 1 {
		number=3*number+1
		values.Add(number)
	}
	if number==1{
		values.Add(1)
	}else{
		Collatz(number,values,false)
	}
}