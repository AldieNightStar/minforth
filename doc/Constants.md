# Constants

## Notes
* Constants are compile-time variables (Which cannot be directly used by user)
* They are used to control blocks and other game stuff
* Constants are needed for only call the functions
* You can't reuse values from constants
* Note that once you set constant it will last until the next set. Be careful


## Set new constant
* To set new constant - just do `:name=value`
* Here are sample how to use `control` command
    * `control` command is using `p` and `b` constants as `parameter` and `block`
```
:p=enabled :b=block1 1 control
```