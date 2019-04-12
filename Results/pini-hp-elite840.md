<img src="../Assets/pini-hp-elitebook-840.png" alt="home spec" width="350px"/>

# CALCULATION (cpu intensive)
|FW && Language | Measure | Best Result| Comments
|----|-|-|-|
|nodejs v10.15.1 - javascript   | ms    |-  | the first time took longer
|.Net Core v2.2 - C#            | ms    |- | no suprise for me node js is weak on memory ops
| GO v1.12                      | ms    |-  | a real big suprise for me (this lang should preform as c/c++)

# File System (io intensive)
|FW && Language | Create Files | Write To File | Best Result| Comments
|----|-|-|-|-|
|nodejs v10.15.1 - javascript   | ms   | ms  | - | no suprise this what is famous for
|.Net Core v2.2 - C#            | ms | ms |-  |since run on the same core as nodejs (ported to c# I thought the numbers will be much closer) **writing in parallel did not worked**
| GO v1.12                      | ms |ms |-  | DISAPPOINTING in some cases not all files were created (althought could be my fault) 

# HTTP (network intensive)
|FW && Language | GET | POST | Best Result| Comments
|----|-|-|-|-|
|nodejs v10.15.1 - javascript   | ms    | ms    | - | no suprise this what is famous for
|.Net Core v2.2 - C#            | ms    | ms    |-  | need to see if the result are linear or not
| GO v1.12                      | ms    | ms    |-  | 