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

## Thoughts on Go
I like strong static typing. I like it _alot_.

I like it because it because it requires the developer to think of all possible scenarios, explicitly limit the scope of the code, erasing unknowns.

Gos type system is small. While small may be easy to grasp, it also means there are plenty of things that are hard to express, _alot harder than in other languages_ - and for no apparent reason. For instance - Why are there no proper enums in Go?

This lack of expressive types is also visible in method return types. In the [tour](https://go.dev/tour/basics/6) one can read:
>  A function can return any number of results

Multiple results? Really? Isn't that a single result of type (string, string)? Instead of allowing that type to be expressed clearly, it's as if the language authors is trying to hide it from the developer. Why?
