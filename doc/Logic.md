# Logic

## Note
* Logic needed to check some conditions
* If conditions are met then we doing a jump
* If not, then no jumps happens
* Condition is any type of comparing after which we have [Jump token](Labels.md)


## Condition types
* Each conditions takes last two values from the stack and doing comparation
    * `<` - less than second the value
    * `>` - greater than the second value
    * `<=` - less or equal to the second value
    * `>=` - greater or equal to the second value
    * `=` - equal to the second value
    * `!=` - not equal to the second value


## Condition Usage
* If `2` is equal to `2` then jump to `start` label
```
2 2 = start!
```
* If `3 2 +` is equal to `5` then jump to `start` label
```
3 2 + 5 = start!
```
* If `a` is greater than `20` then jump to `start` label
```
$a 20 > start!
```
* If `a` is less than `0` then jump to `start` label
```
$a 0 < start!
```