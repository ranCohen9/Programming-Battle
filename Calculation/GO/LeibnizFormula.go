package main
import (
	"math"
	"fmt"
	"time"
)

func hello()(string){
	return "hello from formula";
}

func SeriesCalculation(){
	var seriesSum float64=0;
	goal :=math.Pi/4;
	fmt.Println("this is the goal ",goal);
	start :=time.Now();
	n:=0
	for ; math.Abs(seriesSum-goal)> 0.000001 && n < 300000 ; n++ {
		seriesSum+=LeibnizOnce(float64(n));
	}
	stop:=time.Now();
	fmt.Println("elapsed: ", stop.Sub(start), "| score: ",seriesSum, "({",n,"})");
}

func LeibnizOnce(n float64)(float64){
	return ((math.Pow(-1,n))/(2*n+1));
}