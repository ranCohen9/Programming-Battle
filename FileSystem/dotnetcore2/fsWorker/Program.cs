using System;

namespace fsWorker
{
    class Program
    {
        static void Main(string[] args)
        {
            FsWorker.FsWorker fsWorker = new FsWorker.FsWorker();
            fsWorker.CreateManyFiles();
            fsWorker.WriteToFile();
        }
    }
}
