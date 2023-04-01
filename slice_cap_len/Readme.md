### Capacity of slice
Capacity of a slice is the number of element in the underlying array, but counting from the *first element* of the slice.
```
// len=9, cap=9
a := [...]int{1,2,3,4,5,6,7,8,9}
// len=1, cap=9, because the first element of the slice is 1 and last element of array is 9
// if we count from 1..9 we get 9 element  
s := [:1]
// len=0, cap=9, zero lenth but capacity still is 9
s := [:0]
// len=1, cap=1, because the first element of slice is 9 and last element of array is also 9
s := [8:]
```
Capacity of a slice is counted from the first element of the slice.
