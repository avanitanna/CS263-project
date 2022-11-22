## Experimentation

We ran each benchmark 4 times and report the average. For experimentation we used the Unix time command that provided metrics like - elapsed time, Max Resident Set Size, CPU utlized, page faults etc.Table 2 summarizes our findings which is later discussed in the closing of section 3. Matrix Multiplication, Channels/Queues and Logging use threads in the range of 5 to 10. Our code for the programs can be found here [https://github.com/avanitanna/CS263-project](https://github.com/avanitanna/CS263-project) . For Go we specifically used the language version “go1.19.3 linux”, and Python 3.10.6 (CPython).

---

## Results Summary

| Application | Python | Go |
| --- | --- | --- |
| Maximum resident set size (kbytes) | Elapsed (wall clock) time (m:ss) | Percent of CPU this job got | Minor (reclaiming a frame) page faults | Maximum resident set size (kbytes) | Elapsed (wall clock) time (m:ss) | Percent of CPU this job got | Minor (reclaiming a frame) page faults |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Matrix Multiplication | 29536 | 1:19.71 | 100 | 7877 | 24628 | 0:03.90 | 217 | 30653 |
| File operations | 6796 | 0:05.57 | 11 | 1735 | 33076 | 0:07.70 | 14 | 30465 |
| Channels/Queues | 7120 | 8:19.28 | 0 | 1821 | 37536 | 8:10.15 | 0 | 31168 |
| Graph BFS | 267284 | 0:05.93 | 98 | 67027 | 364636 | 0:03.01 | 104 | 118901 |
| Merge Sort | 101632 | 1:03.50 | 101 | 25938 | 438828 | 0:01.82 | 113 | 142966 |
| Logging  | 8292 | 0:08.63 | 1 | 2169 | 38168 | 0:07.78 | 4 | 31305 |

                         Table 2: Our applications and runtime performance of Python and Go