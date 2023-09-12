package main

import (
	"fmt"
	"os"

	//"log"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/wcharczuk/go-chart/v2"
)

func main(){
	var number int
	fmt.Print("Enter a integer number for Collatz Conjecture:")
	fmt.Scan(&number)
	
	ycoords,xcoords:=assign(number)
	titlestr:=fmt.Sprintf("Collatz Conjecture chart for %d", number)

	graph := chart.Chart{
		Title: titlestr,
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xcoords,
				YValues: ycoords,
			},
		},
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
	
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
func assign(number int)([]float64,[]float64){
	var yv interface{}
	values:=arraylist.New()
	Collatz(number,values,true)
	ycoordinates:= []float64{}
	xcoordinates:=[]float64{}

		for i:=0 ;i<values.Size();i++{
			yv,_=values.Get(i)
			ycoordinates = append(ycoordinates, float64(yv.(int)))
			xcoordinates=append(xcoordinates, float64(i+1))	
		}

	
	return ycoordinates,xcoordinates
}