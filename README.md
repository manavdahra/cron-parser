Cron Expression parser
==

This utility enables a user to pass in arguments as a combination of `cron expression + command` in standard format
`(Example: */15 0 1,15 * 1-5 /usr/bin/find)`
and spits out the output in exploded format as follows:
```text
minute        0,15,30,45
hour          0
day of month  1,15
month         1,2,3,4,5,6,7,8,9,10,11,12
day of week   1,2,3,4,5
command       /usr/bin/find
```

Build & Run:
==
To run the utility follow the below mentioned steps:

1. `make all` (under root project folder)
   
    Depending upon different system architectures and platforms, binary files shall be generated under the path `bin` folder
   
2. `bin/deliveroo-cron-darwin-amd64 "* 3 1-10 */3 6 /usr/bin/find"`
    ```text
    minute        0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59
    hour          3
    day of month  1,2,3,4,5,6,7,8,9,10
    month         3,6,9,12
    day of week   6
    command       /usr/bin/find
    ```
   
![Example](ss.png "Steps")

Unhandled fields/expressions
==

For the sake of simplicity, only standard CRON expression format is supported for now

List of **not** supported fields:
```text
@yearly
@annually
@monthly
@weekly
@daily
@hourly
@reboot
L 
W
#
?
H
```
