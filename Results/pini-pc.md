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
|nodejs v10.15.1 - javascript   | 0.714ms    | 0.742ms | X    | no suprise this what is famous for
|.Net Core v2.2 - C#            | 3.23ms    | 5.66ms |-     | COULD NOT PREFORM THE 1000 iterations server threw excetions so I did 100*100 need to see if the result are linear or not
| GO v1.12                      | 0.66ms       |  0.71   |-     | the post testing threw errors on queue full and lack of space even after I reduced the iterations to 100*100 so this results are a little faulty. the GET part is well preformed, but combing this results with the rest got to make you think