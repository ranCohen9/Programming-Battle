using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace client
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("running http battle");
            Worker t=new Worker();
            //t.GetMany().Wait();
            t.AvgGetMany().Wait();
            t.PostMany().Wait();
            System.Console.WriteLine("client exited");
        }
    }
}
