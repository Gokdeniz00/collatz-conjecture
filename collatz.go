package main

import (
	"fmt"
	"log"

	"github.com/emirpasic/gods/lists/arraylist"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func main(){
	var number int
	fmt.Print("Enter a number for Collatz Conjecture:")
	fmt.Scan(&number)
	values := arraylist.New()
	Collatz(number, values,true)
	fmt.Print(values.Get(0))

	p:=plot.New()
	
	p.Title.Text = "Collatz Conjecture"
	p.X.Label.Text="Step"
	p.Y.Label.Text="Value"

	p.Add(plotter.NewGrid())
	points:=Points(values.Size(),values)

	err := plotutil.AddLinePoints(p,points)
	if err != nil{
		log.Fatal(err)
	}

}
func Points(size int, values *arraylist.List) plotter.XYs {
	pts:=make(plotter.XYs, size)
	for i:= range pts{
		pts[i].X=float64(i+1)
		y,_:=values.Get(i)
		pts[i].Y=float64(y)
	}
	return pts
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