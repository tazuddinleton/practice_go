`defer` pushes a function call onto a stack
### 3 rules of `defer`

1. The argument of `defer` is evaluated when `defer` is evaluated
```
func a() {
    i := 0
    defer fmt.Println(i)
    i++
}
```
this will print '0' when called, not '1' because when `defer` was encountered value of i was 0.
2. `defer`ed functions are called in Last In First Out order
```
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Printf("%d",i)
    }
}
```
this will print '3210' in the console when called.
3. `defer`ed functions may read and assign to the named return values of surrounding function
```
func c() (i int) {
    defer func(){i++}()
    return 1
}
```
this will actually return 2 because `defer`ed func will modify the return value.


### panic
`panic` takes control from normal execution flow and starts unwinding the stack. `defer`ed functions are
executed normally during panic. 
```
func g(i int) {
    if i > 3 {
        fmt.Println("panicinig!")
        panic(i)
    }
    fmt.Println("incrementing...", i)
    g(i+1)
}

```
`panic`ed function returns to its caller and the caller

### recover
`recover` recovers from a panic and gets back to normal execution flow. It only useful in `defer`ed func.
Calling recover in normal flow does nothing but returning `nil`. 

```
func f() {
    fmt.Println("calling g")
    g(0)
}

func main() {
    f()
    fmt.Println("returned normally from f")
}
```
calling above would output the following
```
> calling g
> incrementing 0
> incrementing 1
> incrementing 2
> incrementing 3
> panicing!
> panic 4
```
but if we recover in `f` like below
```
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered in f", r)
        }
    }()
    fmt.Println("calling g")
    g(0)
}
```
then the output will be following
```
> calling g
> incrementing 0
> incrementing 1
> incrementing 2
> incrementing 3
> panicing!
> recovered in f 4
> returned normally from f
```

For detail [see this article](https://go.dev/blog/defer-panic-and-recover)