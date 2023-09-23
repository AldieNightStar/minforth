# Control blocks

## Note
* To control the block you need two constants:
    * `b` - block which will be controlled. For example: `block1`
    * `p` - parameter of the block. For example `enabled`
* Then `control` accepts some single stack value like `1`


## Usage
```
:b=block1 :p=enabled 1 control
```

## Sample: Turn off/on the conveyor
```
:b=conveyor1 :p=enabled 0 control
10 wait
:b=conveyor1 :p=enabled 1 control
```