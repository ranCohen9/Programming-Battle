<img src="../Assets/homepc.png" alt="home spec" width="350px"/>

# CALCULATION (cpu intensive)
|FW && Language | Measure | Best Result| Comments
|----|-|-|-|
|nodejs v10.15.1 - javascript   | 15ms  |-  | the first time took longer
|.Net Core v2.2 - C#            | 5.5ms | X | no suprise for me node js is weak on memory ops
| GO v1.12                      | 22ms  |-  | a real big suprise for me (this lang should preform as c/c++)

# File System (io intensive)
|FW && Language | Create Files | Write To File | Best Result| Comments
|----|-|-|-|-|
|nodejs v10.15.1 - javascript   | 9ms   | 13ms  | X | no suprise this what is famous for
|.Net Core v2.2 - C#            | 133ms | 224ms |-  |since run on the same core as nodejs (ported to c# I thought the numbers will be much closer) **writing in parallel did not worked**
| GO v1.12                      | 321ms |3500ms |-  | DISAPPOINTING in some cases not all files were created (althought could be my fault) 

# HTTP (network intensive)
|FW && Language | GET | POST | Best Result| Comments
|----|-|-|-|-|
|nodejs v10.15.1 - javascript   | 1.59ms    | 1.41ms | X    | no suprise this what is famous for
|.Net Core v2.2 - C#            | 3.31ms    | 2.54ms |-     | need to see if the result are linear or not
| GO v1.12                      | TBD       |  TBD   |-     | 