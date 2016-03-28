# hap

Command that makes the haproxy stats socket more useful.

```shell
$ hap --help
  -info
        Prints out helpful definitions
  -socket string
        Haproxy stats socket to connect to (default "/var/run/haproxy")

$ sudo hap -info -socket /var/run/haproxy.sock
pxname: Proxy Name
svname: Service Name
qcur: Current Queue
qmax: Max Queue
scur: Current Sessions
smax: Max Sessions
slim: Session Limit
stot: Cumulative Connections
bin: Bytes In
bout: Bytes Out
ereq: Request Errors
econ: Connection Errors
eresp: Response Errors
status: Status
rate: Request Rate
qtime: Avg Queue Time
ctime: Avg Connect Time
rtime: Avg Response Time
+-----------------+-----------------+------+------+------+------+------+--------+------------+------------+------+------+-------+----------+------+-------+-------+-------+
|      PXNAME     |     SVNAME      | QCUR | QMAX | SCUR | SMAX | SLIM |  STOT  |    BIN     |    BOUT    | EREQ | ECON | ERESP |  STATUS  | RATE | QTIME | CTIME | RTIME |
+-----------------+-----------------+------+------+------+------+------+--------+------------+------------+------+------+-------+----------+------+-------+-------+-------+
| server-01       | FRONTEND        |      |      |    0 |  125 | 2000 |  18438 |    1009187 |     979916 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-01       | server-01       |    0 |    0 |    0 |  125 |      |  18832 |    1009187 |     979916 |      |   66 |     0 | no check |    0 |     0 |    77 |     0 |
| server-01       | BACKEND         |    0 |    0 |    0 |  125 |  200 |  18438 |    1009187 |     979916 |      |   66 |     0 | UP       |    0 |     0 |    77 |     0 |
| server-02       | FRONTEND        |      |      |    0 |  144 | 2000 |  18427 |    1008606 |     974616 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-02       | server-02       |    0 |    0 |    0 |  144 |      |  18665 |    1008606 |     974616 |      |   12 |     0 | no check |    0 |     0 |    43 |     0 |
| server-02       | BACKEND         |    0 |    0 |    0 |  144 |  200 |  18427 |    1008606 |     974616 |      |   12 |     0 | UP       |    0 |     0 |    43 |     0 |
| server-03       | FRONTEND        |      |      |    0 |  149 | 2000 |  18027 |     986606 |     955751 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-03       | server-03       |    0 |    0 |    0 |  149 |      |  18372 |     986606 |     955751 |      |   43 |     0 | no check |    0 |     0 |   142 |     0 |
| server-03       | BACKEND         |    0 |    0 |    0 |  149 |  200 |  18027 |     986606 |     955751 |      |   43 |     0 | UP       |    0 |     0 |   142 |     0 |
| server-04       | FRONTEND        |      |      |    0 |  113 | 2000 |  17764 |     972141 |     954101 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-04       | server-04       |    0 |    0 |    0 |  113 |      |  17972 |     972141 |     954101 |      |   16 |     0 | no check |    0 |     0 |    53 |     0 |
| server-04       | BACKEND         |    0 |    0 |    0 |  113 |  200 |  17764 |     972141 |     954101 |      |   16 |     0 | UP       |    0 |     0 |    53 |     0 |
| server-05       | FRONTEND        |      |      |    0 |  135 | 2000 |  17601 |     963176 |     927921 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-05       | server-05       |    0 |    0 |    0 |  135 |      |  17968 |     963176 |     927921 |      |   58 |     0 | no check |    0 |     0 |    48 |     0 |
| server-05       | BACKEND         |    0 |    0 |    0 |  135 |  200 |  17601 |     963176 |     927921 |      |   58 |     0 | UP       |    0 |     0 |    48 |     0 |
| server-06       | FRONTEND        |      |      |    0 |  144 | 2000 |  17601 |     963104 |     931790 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-06       | server-06       |    0 |    0 |    0 |  144 |      |  17904 |     963104 |     931790 |      |   34 |     0 | no check |    0 |     0 |    76 |     0 |
| server-06       | BACKEND         |    0 |    0 |    0 |  144 |  200 |  17601 |     963104 |     931790 |      |   34 |     0 | UP       |    0 |     0 |    76 |     0 |
| server-07       | FRONTEND        |      |      |    0 |  116 | 2000 |  17764 |     972028 |     943164 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-07       | server-07       |    0 |    0 |    0 |  116 |      |  17992 |     972028 |     943164 |      |   21 |     0 | no check |    0 |     0 |    51 |     0 |
| server-07       | BACKEND         |    0 |    0 |    0 |  116 |  200 |  17764 |     972028 |     943164 |      |   21 |     0 | UP       |    0 |     0 |    51 |     0 |
| server-08       | FRONTEND        |      |      |    0 |  139 | 2000 |  17656 |     965064 |     929149 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-08       | server-08       |    0 |    0 |    0 |  139 |      |  18094 |     965064 |     929149 |      |   56 |     0 | no check |    0 |     0 |   209 |     0 |
| server-08       | BACKEND         |    0 |    0 |    0 |  139 |  200 |  17656 |     965064 |     929149 |      |   56 |     0 | UP       |    0 |     0 |   209 |     0 |
| server-01_thing | FRONTEND        |      |      |    0 |   32 | 2000 |  23628 |    1295471 |    9667047 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-01_thing | server-01_thing |    0 |    0 |    0 |   32 |      |  24308 |    1295471 |    9667047 |      |  104 |     0 | no check |    0 |     0 |    91 |     0 |
| server-01_thing | BACKEND         |    0 |    0 |    0 |   32 |  200 |  23628 |    1295471 |    9667047 |      |  104 |     0 | UP       |    0 |     0 |    91 |     0 |
| server-02_thing | FRONTEND        |      |      |    0 |   36 | 2000 |  23641 |    1296186 |    9679053 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-02_thing | server-02_thing |    0 |    0 |    0 |   36 |      |  24189 |    1296186 |    9679053 |      |   88 |     0 | no check |    0 |     0 |    73 |     0 |
| server-02_thing | BACKEND         |    0 |    0 |    0 |   36 |  200 |  23641 |    1296186 |    9679053 |      |   88 |     0 | UP       |    0 |     0 |    73 |     0 |
| server-03_thing | FRONTEND        |      |      |    0 |   34 | 2000 |  23621 |    1295086 |    9673671 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-03_thing | server-03_thing |    0 |    0 |    0 |   34 |      |  24185 |    1295086 |    9673671 |      |   81 |     0 | no check |    0 |     0 |   127 |     0 |
| server-03_thing | BACKEND         |    0 |    0 |    0 |   34 |  200 |  23621 |    1295086 |    9673671 |      |   81 |     0 | UP       |    0 |     0 |   127 |     0 |
| server-04_thing | FRONTEND        |      |      |    0 |   29 | 2000 |  23545 |    1290906 |    9640965 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-04_thing | server-04_thing |    0 |    0 |    0 |   29 |      |  24105 |    1290906 |    9640965 |      |   84 |     0 | no check |    0 |     0 |    95 |     0 |
| server-04_thing | BACKEND         |    0 |    0 |    0 |   29 |  200 |  23545 |    1290906 |    9640965 |      |   84 |     0 | UP       |    0 |     0 |    95 |     0 |
| server-05_thing | FRONTEND        |      |      |    0 |   33 | 2000 |  23379 |    1281776 |    9573483 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-05_thing | server-05_thing |    0 |    0 |    0 |   33 |      |  23896 |    1281776 |    9573483 |      |   81 |     0 | no check |    0 |     0 |   109 |     0 |
| server-05_thing | BACKEND         |    0 |    0 |    0 |   33 |  200 |  23379 |    1281776 |    9573483 |      |   81 |     0 | UP       |    0 |     0 |   109 |     0 |
| server-06_thing | FRONTEND        |      |      |    0 |   34 | 2000 |  23463 |    1286396 |    9611571 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-06_thing | server-06_thing |    0 |    0 |    0 |   34 |      |  23987 |    1286396 |    9611571 |      |   73 |     0 | no check |    0 |     0 |   105 |     0 |
| server-06_thing | BACKEND         |    0 |    0 |    0 |   34 |  200 |  23463 |    1286396 |    9611571 |      |   73 |     0 | UP       |    0 |     0 |   105 |     0 |
| server-07_thing | FRONTEND        |      |      |    0 |   32 | 2000 |  23369 |    1281226 |    9573483 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-07_thing | server-07_thing |    0 |    0 |    0 |   32 |      |  23862 |    1281226 |    9573483 |      |   71 |     0 | no check |    0 |     0 |    61 |     0 |
| server-07_thing | BACKEND         |    0 |    0 |    0 |   32 |  200 |  23369 |    1281226 |    9573483 |      |   71 |     0 | UP       |    0 |     0 |    61 |     0 |
| server-08_thing | FRONTEND        |      |      |    0 |   38 | 2000 |  23414 |    1283701 |    9581763 |    0 |      |       | OPEN     |    0 |       |       |       |
| server-08_thing | server-08_thing |    0 |    0 |    0 |   38 |      |  23987 |    1283701 |    9581763 |      |   96 |     0 | no check |    0 |     0 |    86 |     0 |
| server-08_thing | BACKEND         |    0 |    0 |    0 |   38 |  200 |  23414 |    1283701 |    9581763 |      |   96 |     0 | UP       |    0 |     0 |    86 |     0 |
| queue           | FRONTEND        |      |      |  152 |  185 | 2000 |    416 |  385633031 |  755390465 |    0 |      |       | OPEN     |    0 |       |       |       |
| queue           | my-queue-01     |    0 |    0 |   51 |   62 |      |    139 |  128950786 |  253110274 |      |    0 |     0 | UP       |    0 |     0 |     1 |     0 |
| queue           | my-queue-02     |    0 |    0 |   50 |   62 |      |    135 |  154034310 |  251472926 |      |    0 |     0 | UP       |    0 |     0 |     0 |     0 |
| queue           | my-queue-03     |    0 |    0 |   51 |   61 |      |    142 |  102647935 |  250807265 |      |    0 |     0 | UP       |    0 |     0 |     0 |     0 |
| queue           | BACKEND         |    0 |    0 |  152 |  185 |  200 |    416 |  385633031 |  755390465 |      |    0 |     0 | UP       |    0 |     0 |     1 |     0 |
| db              | FRONTEND        |      |      |   16 |  363 | 2000 | 184731 | 1152123486 | 2666449779 |    0 |      |       | OPEN     |    0 |       |       |       |
| db              | my-db-01        |    0 |    0 |    5 |  121 |  500 |  61680 |  391094012 |  904371197 |      |   17 |     0 | UP       |    0 |     0 |     0 |     0 |
| db              | my-db-02        |    0 |    0 |    5 |  121 |  500 |  59753 |  384886771 |  880366479 |      |   12 |     0 | UP       |    0 |     0 |     1 |     0 |
| db              | my-db-03        |    0 |    0 |    6 |  121 |  500 |  63518 |  376142703 |  881712103 |      |    7 |     0 | UP       |    0 |     0 |     0 |     0 |
| db              | BACKEND         |    0 |    0 |   16 |  363 |  200 | 184731 | 1152123486 | 2666449779 |      |   46 |     0 | UP       |    0 |     0 |     1 |     0 |
| api             | FRONTEND        |      |      |    0 |    6 | 2000 |    776 |    8245032 |   26235831 |    9 |      |       | OPEN     |    0 |       |       |       |
| api             | my-api-01       |    0 |    0 |    0 |    3 |      |   5405 |    2758511 |    8761416 |      |    7 |     0 | UP       |    0 |     0 |    70 |   205 |
| api             | my-api-02       |    0 |    0 |    0 |    3 |      |   5366 |    2734540 |    8752742 |      |    5 |     0 | UP       |    0 |     0 |    73 |   202 |
| api             | my-api-03       |    0 |    0 |    0 |    4 |      |   5399 |    2749353 |    8719329 |      |    9 |     0 | UP       |    0 |     0 |    78 |   201 |
| api             | BACKEND         |    0 |    0 |    0 |    6 |  200 |  15994 |    8245032 |   26235831 |      |   25 |     0 | UP       |    0 |     0 |    95 |   179 |
+-----------------+-----------------+------+------+------+------+------+--------+------------+------------+------+------+-------+----------+------+-------+-------+-------+
```