# Math

## Note
* Math operators take last two numbers from stack and do some calculations
* After calculation is made the result is put back to stack

## Operators
* There are four operators:
```
1  1 +
2  1 -
5  2 *
10 2 /
```

## Sample
* Result of this one will be: `8`
    * We added `2 2 +` and instead of `2 2` in stack now only `4`
    * We multiplied `4` by `2 *` and thus we have `8`
```
2 2 +
2 *
```


## Multiple increments
* Result of this one will be `6`
```
1 1 +
1 +
1 +
1 +
1 +
```