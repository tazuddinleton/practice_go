## Array
Array have fixed length. Its length is part of its definition. `[4]int` denotes an integer array of length for, similarly
`[5]int` denotes an array of length 5 whose type is integer. But both are completely distinct type.
```



func doSomething(nums [5]int) {
    
}

```

The above function expects an array with length 5 hence `[4]int` is not compatible with that func.

## Slice
A slice is an abstraction over an array. The type specification for a slice is []T. Where T is any type. Unlike an array
slice has no length specified with its declaration.

Major difference between the two is:
*You have to specify the length of an array when declaring it, whereas you do not specify the length of a slice.*
