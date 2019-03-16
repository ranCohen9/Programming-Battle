using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace client
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!2");
            Worker t=new Worker();
            t.GetMany().Wait();
            t.PostMany().Wait();
            System.Console.WriteLine("client exited");
        }
    }
}
