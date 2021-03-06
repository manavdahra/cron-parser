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
> Note: for the sake of simplicity, only standard CRON expression format is supported for now

Following fields are **not** supported yet
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


