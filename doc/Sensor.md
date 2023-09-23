# Sensor

## Notes
* Used to take information from specific block
* After information took then it's pushed to the stack
* Uses constants `b` and `p` which is `block` and `parameter`


## Samples
* Let's check how many copper `block1` has
```
:b=block1 :p=@copper sensor
```
* Let's check how many energy battery has
```
:b=block1 :p=@totalPower sensor
```


## Logic check
* Each iteration will wait for `10` seconds to not overload
* Let's turn off the conveyor if `coopper` is more than `1000`
* And start again when less than `200`
```
loop:
    10 wait
    :b=block1 :p=@copper sensor 1000 > stop!
    :b=block1 :p=@copper sensor  200 < start_again!
    loop!
stop:
    :b=conveyor1 :p=enabled 0 control
    loop!
start_again:
    :b=conveyor1 :p=enabled 0 control
    loop!
```