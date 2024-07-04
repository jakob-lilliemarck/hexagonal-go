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
I like strong static typechecking at compile time. I like it _alot_.

I like it because it because it requires the developer to think of all possible scenarios, explicitly limit the scope of the code, erasing unknowns - and it prohibits running the code before it has a fair chance working.

Why would you ever run code before you know it's working anyway? Languages that allow you to, tend to pull you into an almost infintate turmoil of _guessing and hoping_. I'm much happier when I'm at _knowing and delivering_.

Gos type system is small. Small is good as it typically means easy to grasp and quick to learn. But small can also mean less expressivity. In other words, there things that are hard to express, _alot harder than in other languages_, and for no apparent reason. For instance, why are there no enums in Go?

This lack of expressive types is also visible in method return types. In the [tour](https://go.dev/tour/basics/6) one can read:
>  A function can return any number of results

Multiple results? Really? Isn't that a single result of type (string, string)? Instead of allowing that type to be expressed clearly, it's as if the language authors is trying to hide it from the developer. Why?

