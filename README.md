# Hexagonal-go
Example RESTful web API using the ports and adapters pattern in Go

Build
```console

```

Run 
```
./main
```

## Thoughts on Go
These are the things I miss in Go
1. Expressive type system (Enums, method generics, ..)
2. Explicit interfaces
3. Elegant error handling
4. Move-semantics

### Expressive type system
I like strong static typechecking at compile time. I like it _alot_.

I like it because it because it requires the developer to think of all possible scenarios, explicitly limit the scope of the program and elimitating unknowns - and it prohibits running the code before the code actually has a fair chance at doing what you expect it to do. Why would you ever run code before that point anyway? Languages that allow you to do so, tend to pull you into an almost infintate turmoil of _guessing and hoping_. As a developer, thats my sad place.

Gos type system is small, and small can be really good! It typically means easy to grasp and quick to learn. But small can also mean less expressivity and that some things become harder to represent as one might otherwise do. It also means some things are a lot _alot harder_ to represent than in other languages. For instance, why are there no enums in Go?

### Explicit interfaces
One of the things I really like with strongly typed code, is the saftey it provides while refactoring. Change a method signature in an interface, and you'll immediately see your file-tree view light up in red indicating where you need to look and refactor to support the new method signature.

To me, _the whole point_ of compile-time errors and strong typing is to support the developer in avoiding unnecessary, often trivial even, bugs. Thinking about that for a moment - that the supporting the developer is the ultimate goal of typing - it means it's also of great importance _how those errors are presented to the developer_. Are the errors displayed close to where the code needs refactoring? Are the error messages understandable? ...and so on.

Go has quite good compiler messages. When an interface is not met it points out which signature mismatches clearly stating that it got and what it expected to get. However, I think it throws the error in the wrong place.

Go implements interfaces implicitly. That means the structs themselves don't know what interfaces they implement, only the caller knows if the struct satisfies the requirement. As such Go opts to display the error in the caller, not in the struct or method implementation.

I typically find myself refactoring the interfaces prior refactoring its implementors. Think about it, it makes little sense to do it the other way around, how would you know what to change the implementor into if the interface has not changed? Changing the interface first means I then would like the compiler to point me to the implementors that's are supposed to implement the new contract. Instead, Go points me to the caller, but the caller has done nothing wrong!

### Elegant error handling
Errors should always be handled exhaustivly. There should be no unhandled errors in a program.

"Handling" however could mean anything, but whatever it is, it should always be done explicitly within the current code base. Letting libraries throw errors for you probably means you're not aware of all scenarios at each point in your programs execution, and not beeing aware of what your program does is a _bad thing_. 

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

Handling errors explicitly typically means you want to pass them to a function that transform them into a new type. It could be a new type of error, or some fallback value. What it is doesn't really matter, in generic terms the signature would be the same:

```
T => U
```

But wait, there's more to it. Functions return either the desired value _or_ they return an error (`nil` is _not_ a value). In other words, they return a single value that is one of two types. That is an enum, tagged union, whatever you want to call it!

So falliable functions could instead return an enum `R` with two members that hold generic values, `T` for the success type and `E` for the error type. Handling an error from such a function could then be described as:
```
R<T, E> =>  U
```

Or in the case of translating result type, for instance when writing layered applications:
```
R<T, E1> => R<T, E2>
```

That looks like candiate for a monad, does it not? And sure, I could write my own result monads, [I actually did](src/lib/utils/result.go) and it works, but due to the limitations of Go generics it just doesn't get very nice and clean anyway. That's because in Go, _methods can't take generic arguments_. Method generics must be defined on the implementing struct, which means that the struct must make assumptions about how a consumer of it might want to use the return values from its methods. Should it allow mapping to a single error? ..or into to two different errors perhaps? ..and what about the success values, should those be mapped? Walking down that lane will not end good. 

### Move semantics
