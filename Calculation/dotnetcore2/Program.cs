using System;
namespace ShowDown
{
    class Program
    {
        static void Main(string[] args)
        {
            LeibnizFormula liebinz=new LeibnizFormula();
            //Console.WriteLine(liebinz.LeibnizOnce(0));
            liebinz.SeriesCalculation();
        }
    }
}
