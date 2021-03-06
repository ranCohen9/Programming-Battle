using System;
class LeibnizFormula
{

    public void AvgSeriesCalculation(){
        double sum=0;
        int i=0;
        for (; i < 1000; i++)
        {
            sum+= this.SeriesCalculation();
        }
        Console.WriteLine($"\nTotal Avegrage of {i} iterations: {sum/i}ms");
    }
    
    public double SeriesCalculation(){
        double seriesSum=0;
        double goal=Math.PI/4;
        Console.WriteLine($"this is the goal {goal}");
        DateTime start=DateTime.Now;
        int n=0;
        for (; Math.Abs(seriesSum-goal) >0.000001 &&  n<300000 ; n++)
        {
            seriesSum+=this.LeibnizOnce(n);
        }
        DateTime stop=DateTime.Now;
        Console.WriteLine($"elapsed: {(stop-start).TotalMilliseconds} | score: {seriesSum} ({n})");
        return (stop-start).TotalMilliseconds;
    }
    public double LeibnizOnce(int n){
        return (Math.Pow(-1,n))/(2*n+1);
    }
}