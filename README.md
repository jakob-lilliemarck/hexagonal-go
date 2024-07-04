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

I like it because it because it requires the developer to think of all possible scenarios, explicitly limit the scope of the program and elimitating unknowns - and it prohibits running the code before the code actually has a fair chance at doing what you expect it to.

Why would you ever run code before you can expect it to work anyway? Languages that allow you to do so, tend to pull you into an almost infintate turmoil of _guessing and hoping_. I think that's a terrible place to be.

Gos type system is small. Small good. It typically means easy to grasp and quick to learn. But small can also mean less expressivity and that some things become harder to represent as one might otherwise do. Sometimes that means they're _alot harder_ to represent than in other languages. For instance, why are there no enums in Go?

### Explicit interfaces
One of the things I really like with strongly typed code, is the saftey they provide while refactoring. Change a method signature in an interface, and you'll immediately see your tree-view light up in red, indicating where you need to look and refactor to support the new signature.

To me, the whole point of compile-time errors and strong typing is support the developer avoid unnecessary bugs. Thinking about that for a moment - that the supporting the developer is the goal - it means it's also of significance _how those errors are presented_, as that too will impact the developers effectiveness at resolving them.

Go has very good compiler messages. When an interface it's not met, it points our which signature mismatches. However, I think it throws the error in the wrong place.

Go implements interfaces implicitly. That means the structs themselves don't know themselves if they're implementing an interface or not, only the caller knows. As the struct don't know about that Go opts to display the error in the caller.

I typically find myself refactoring the interfaces prior refactoring their implementors. Think about it, it makes little sense to do it the other way around, how would you know what to change the implementor into? That means I would like the compiler to point me to the implementor that's drifted away from the contract, rather than to the caller. The caller has done nothing wrong.

### Elegant error handling

### Move semantics
