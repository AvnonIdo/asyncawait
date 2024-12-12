# AsyncAwait
The library that ports the async/await functionality into Go!
## What does this library do?
 AsyncAwait, as its name suggests, adds the async/await functionality you might be familiar with from other langauges such as Javascript, Rust, C# ect.

 ```go
 import (
    "time"
    "github.com/AvnonIdo/asyncawait"
    )

 func slowFunctionThatTakesTime() int {
    time.Sleep(10)
    return 42
 }

 func main() {
    future := asyncawait.Async(slowFunctionThatTakesTime)
    // future is of type Future[Int]
 }
 ```
 The *async* functions start to run the received function in the background and return a *future*, which can later be *await*-ed in order to obtain the result.

 *Futures* can be *await*-ed by using the `asyncawait.Await` function
 ```go
 num := asyncawait.Await(future)
 // num is of type int
 ```
 Or alternatively by calling `Future.Await`
 ```go
 num := future.Await()
 // num is of type int
 ```
 ## But wait, isn't Go's concurrency model superior to the async/await model?
 Go's concurrency model is indeed excellent, and I implore you to continue to use it to its fullest extent!
 *Channels* and *Select* often offer a much more elegant solution to a concurrency problem than *async* and *await* do.

 However, there are some situations where async/await could cut a LOT of boilerplate, and this is where this library comes in.

 Consider the following code: 
 ```go
 func callSlowFunctionAndMultiplyBySeven() {
    num, err := someSlowFunction()
    if err != nil {
        // error handling
    }

    return num * 7
 }
 ```
 You have some function that takes time; Maybe it fetches something from the internet or reads from a file, and it returns some value and an error.
 You need the result of the function, so you call it, handle the errors if there are any, and use the value. So far so good.

 Now imagine you needed to use another slow function:
 ```go
 func callBothSlowFunctionsAndPrint() {
    num, err := someSlowFunction()
    if err != nil {
        // error handling
    }

    str, err := someOtherSlowFunction()
    if err != nil {
        // error handling
    }

    fmt.Println(num, str)
 }
 ```
Seems simple enough at first, but oh no! The parent function now takes 2 times as long to complete!

The first naive solution is to just prepend a `go`commandc
```go
go someSlowFunction()
go someOtherSlowFunction()
```
Here, done, both functions run at the same time. But wait, what about the return values?!

We need to use channels in order to get the values out of the goroutines.
But channels can't send a pair of values so we need to wrap the `(int,err)` and `(string,err)` in a struct...

Eventually we end up with something like this:
 ```go
 func callBothSlowFunctionsAndPrint() {
    type firstFunctionReturn struct {
        num int
        err error
    }

    ch1 := make(chan firstFunctionReturn)
    go func () {
        num, err := someSlowFunction()
        ch1 <- firstFunctionReturn{num; err}
    }()

    type secondFunctionReturn struct {
        str string
        err error
    }

    ch2 := make(chan secondFunctionReturn)
    go func () {
        str, err := someOtherSlowFunction()
        ch2 <- secondFunctionReturn{str; err}
    }()

    res1 := <- ch1
    res2 := <- ch2

    if res1.err != nil {
        // error handling
    }
    if res2.err != nil {
        // error handling
    }
    fmt.Println(num, str)
 }
 ```
Finally! We can run both functions at the same time and get their result at the end... But it does look kinda ugly... And results in way more lines of code... 

If only we could get some sort of object from the `go` statement that would give us the result like a regular function call but without blocking us from initiating the other goroutines...

With AsyncAwait you could simplify this function like this:
```go
 func callBothSlowFunctionsAndPrint() {
    future1 := asyncawait.Async2(someSlowFunction)
    future2 := asyncawait.Async2(someOtherSlowFunction)

    num, err := future1.Await()
    if err != nil {
        // error handling
    }

    str, err := future2.Await()
    if err != nil {
        // error handling
    }

    fmt.Println(num, str)
 }
 ```

 ## Ok but all examples so far have been for functions without arguments. How do I send arguments?
 Good observation, keen reader! Indeed, all examples so far have been for functions that take no arguments.

 To send arguments to a function, simply wrap it with a closure, like so:
 ```go
 func addButSlow(int num1, int num2) int {
    time.Sleep(10)
    return num1+num2
 }

 func main() {
    future := asyncawait.Async2(func (int,int) int {return addButSlow(num1, num2)})
    // future is of type Future[Int]

    sum := future.Await()
    // sum is of type Int
 }
 ```
 I will admit that Go's verbose closures make it a bit clunky, but it is the only way of calling a function with arbitrary arguments without resorting to reflection, which would remove compile-time type safety and harm performance.
 
 ## The numbers! What do they mean?
 By this point you probably noticed we used both `asyncawait.Async` and `asyncawait.Async2`.

 The number 2 simply indicates the fact that the function we are passing in returns 2 values (and therefore, the *future* should return 2 values).

 In fact we have `asyncawait.Async`, `asyncawait.Async2`, `asyncawait.Async3`, `asyncawait.Async3`, `asyncawait.Async4` and `asyncawait.Async5`. 
 Each corresponds to the respective amount of return values, with `asyncawait.Async` corresponding to a single return value.

 There are also corresponding `asyncawait.Await` functions for all these amounts and a respective `asyncawait.Future` type that implements it's own `Await` function.

 There is also `asyncawait.Async0`, `asyncawait.Future0` and `asyncawait.Await0`, included for completeness' sake, which can be used to wait on a function that doesn't return anything. Although in this case you might be better off waiting in the traditional Gophery ways.

## But what if my function returns 6 or more values? 
If you have a function that returns more than 3 values, you should probably refactor it and use a struct as the return type.

If you have a function that returns more than 5 values, then you are simply a madman and I do not feel obligated to support you in this tomfoolery.

Feel free to try and change my mind on this if you want, but add a PR implementing it yourself if you do; It's a Sisyphean task.
