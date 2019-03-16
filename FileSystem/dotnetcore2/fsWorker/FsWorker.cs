namespace FsWorker
{
    using System;
    using System.IO;
    using System.Threading.Tasks;

    class FsWorker
    {
        const string FilesDirName = "files";
        public void CreateDirectory()
        {
            if (!Directory.Exists(FilesDirName))
            {
                Directory.CreateDirectory(FilesDirName);
            }
            else
            {
                string[] files = Directory.GetFiles(FilesDirName);
                foreach (var file in files)
                {
                    File.Delete(file);
                }
            }
        }

        public void CreateFile(int fileNumber)
        {
            File.Create($"{FilesDirName}\\file-{fileNumber}.txt");
        }

        public void CreateManyFiles()
        {
            this.CreateDirectory();
            DateTime start = DateTime.Now;
            Parallel.For(0, 1000, (index) =>
            {
                this.CreateFile(index);
            });
            DateTime stop = DateTime.Now;
            Console.WriteLine($"create files Elapsed: {(stop - start).TotalMilliseconds}");
        }

        public void WriteToFile()
        {
            if (!Directory.Exists(FilesDirName))
            {
                Directory.CreateDirectory(FilesDirName);
            }
            if (File.Exists("contentFile.txt"))
            {
                File.Delete("contentFile.txt");
            }
            // DONT USE STREAM WRITER WE WANT TO TEST OPEN FILE WRITE AND CLOSE IT
            // Parallel for throwing exceptions on file descriptor being used
            DateTime start = DateTime.Now;
            
                for (int i = 0; i < 1000; i++)
                {
                    File.AppendAllText("contentFile.txt",$"index is {i}\n");
                }
            DateTime stop = DateTime.Now;
            Console.WriteLine($"Write to file elapsed: {(stop - start).TotalMilliseconds}");
        }
    }
}