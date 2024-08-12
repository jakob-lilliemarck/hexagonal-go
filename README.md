# Hexagonal-go
Example RESTful web API using the ports and adapters pattern in Go

Build
```console
go build src/app/main.go
```

Run 
```
./main
```

## Some thoughts on Go
This project is my first attempt at learning Go. I'm by no means an expert in the language, and my views or perceptions at the time of writing this might very well be incorrect. Rather than an excercise in correctness however, my intention is to try out and get a sense for Go. Formulating my thoughts in writing is simply a vehicle for making aware and explicit what would otherwise perhaps only be a hunch, an itch or belief. If there are any factual incorrectness I welcome correction and if you don't share my opinion I would be interested in hearing yours!

I've been curious about Go for a while and I've heard really great things about it, so upon starting this project I had quite high expectations and a sense of really wanting to like the language. While I think Go does a lot of things right, I'm going to focus on the things I struggled with and that I felt like I missed while writing this very small example application. After all, the good stuff doesn't really require much debating ;-).

So, these are the things I miss in Go
1. Expressive type system (Enums, method generics, ..)
2. Explicit interfaces
3. Elegant error handling
4. Move-semantics

### Expressive type system
I like strongly typed code and static typechecking at compile time. I like it _alot_.

I like it because it because it requires the developer to think of all possible scenarios, explicitly limit the scope of the program and elimitating unknowns. It prohibits running the code before the code actually has a fair chance at doing what you expect it to do, and why would you ever run your code before that point anyway? Languages that allow yo to run code that can't be expected to work, tend to pull you into an almost infintate turmoil of _guessing and hoping_. As a developer, that is my sad place.

Gos type system is small. Small is great as it typically means it is easy to grasp and and quick to learn. But small can also mean less expressivity and that some things become harder to represent as one might otherwise do. It also means that some things are a lot _alot harder_ to represent than in other languages. For instance, why are there no enums in Go?

### Explicit interfaces
One of the things I really like with strong typing, is the saftey it provides while refactoring. Change a method signature in some interface and you'll immediately see your tree view light up in red, indicating where you need to look and refactor code.

To me, _the whole point_ of compile time errors and strong typing is to support the developer avoiding bugs and writing quality code. Think about that for a moment:
> The whole point of typing is to support the developer

Doesn't that mean that in addition to typing and type-errors themselves, it's also _of great importance how those errors are presented to the developer_? One might ask:
- Are the errors displayed close to where the developer need to tend to the code?
- Are the error messages understandable and _actionable_?

Go has quite good compile-error messages. When an interface is not met the compiler points out which signature mismatches clearly stating what it got and what it expected to get. However, I think it throws the error in the wrong place. Go implements interfaces implicitly. That means the structs themselves don't know what interfaces they implement, only the caller knows if the struct satisfies their requirement. As such Go opts to display the error in the caller,and not in the implementor.

When refactoring I typically find myself changing the interfaces prior its implementors. After all, it would makes little sense to do it the other way around, if the interface has not changed how would you then know how to change its implementor? Changing the interface first means I then would like the compiler to point me to all of the implementors that are no longer honoring the contract. Go instead points me to the caller, but the caller has done nothing wrong!

### Elegant error handling
Errors should always be handled exhaustivly. There should be no unhandled errors in a program.

"Handling" however could mean anything, but whatever it means, it should always be done _explicitly within the current code base_. Letting libraries throw errors for you probably means you're not aware of all scenarios at each point in your programs execution, and not beeing aware of what your program does is a _bad thing_. 

Go does a few really good things with respect to error handling.
1. Go errors require no special syntax like `raise` or `throw`, instead they're returned like any other value.
2. Go requires the developer to assign all returned values of a function, including errors.
3. Go requires the developer to use all assigned values in some way.

Those requirements guarantee that errors are not forgotten and that they must be explicitly dealt with. That is really great, a lot better than most languages!

What Go doesn't do however, is to provide its developer with _ergonomic utilities_ and _abstractions_ to handle errors simply and elegantly. As such, idiomatic error handling in Go looks something like this:

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
```
Simple, but primitive.

Handling errors explicitly often means that you want to transform them into a new type, atleast if you're not into imperative code style (then shame on you!). It could be a new type of error, a response message or a fallback value. What the new type is doesn't really matter, in generic terms the signature would be the same:

```
T => U
```

But wait, there's more to it. Functions return either the desired value _or_ they return an error. `nil` is _not_ a value. In other words, they return _a single value that is one of two types_. That's an enum!

So generically, falliable functions could instead return an enum `R` with two members that hold generic values, `T` for the success type and `E` for the error type. Handling an error from such a function could then be described as:
```
R<T, E> =>  U
```

Or in the case of translating result type, for instance between layers of a layered application:
```
R<T, E1> => R<T, E2>
```

That looks like candiate for a monad, does it not?

Go doesn't provide any conveniences around this pattern. Sure, I could write my own result monads, [I actually did](src/lib/utils/result.go) and it works, but due to the limitations of Go generics it just doesn't get very nice and clean anyway. That's because in Go, _methods can't take generic arguments_. Method generics must be defined on the implementing struct, which means that the struct must make assumptions about how a consumer of it might want to use the return values from its methods. Should it allow mapping to a single error? ..or into to two different errors perhaps? ..and what about the success values, should those be mapped? Walking down that lane will not end good.

### Move semantics
So finally, "move semantics", what the hell is that even?

Let's start with an example:
```go
package main

import "log"

type payload struct {
	value int
}

func modify(p payload) {
	p.value += 1
}

func main() {
	p := payload{1}
	modify(payload{1})
	log.Println(p.value)
}
```
There's a struct `payload` that holds a single `value` of type `int`, and function `modify(payload)` that takes that same struct as its only argument, mutates its `value` property and doesn't return anything. Now, how many times is memory allocated for `p`? One, you might guess, on the second line of `main()` - but wrong you would be!

Go passes arguments by value. If you would want to pass a pointer you explicitly pass the pointer by value as well. In the example above it means that upon invoking `modify(p)` memory is allocated anew, `p` is copied into that memory space and made available within the function scope. If you run the example program you'll see that `p.value` is still at the time of printing it.

Now there are several things to dislike about that example, including the stupidity of the program itself. However, what I'm trying to illustrate and what threw me off about it, is that the compiler is prefectly happy to let me use `p` after I've passed it _*by value*_ to `modify(p)`. That's because Go doesn't implement any _move semantics_! If it did, the compiler would be able to tell me that there's an error at `log.Println(p.value)` as `p` has been _moved into_ `modify(p)`. 
