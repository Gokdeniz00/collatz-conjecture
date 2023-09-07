package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main(){
	var number int
	fmt.Print("Enter a number for Collatz Conjecture:")
	fmt.Scan(&number)
	values := arraylist.New()
	Collatz(number, values,true)
	fmt.Print(values)

	p:=plot.New()
	
	p.Title.Text = "Collatz Conjecture"
	p.X.Label.Text="Step"
	p.Y.Label.Text="Value"

	p.Add(plotter.NewGrid())
	points:=Points(values.Size(),values.Values())
	//graph,err:=plotter.NewScatter()
}
func Points(size int, values []interface{}) plotter.XYs {
	pts:=make(plotter.XYs, size)

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
	if number!=1{
		Collatz(number,values,false)
	}
}