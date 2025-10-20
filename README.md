# key-value-store
Yet another Key-Value store [With a twist though].

Have a simple app? Don't even need the KVS server running.
Just pipe the values in a local file!

Better even, use an S3 bucket as your KVS.

# New Architecture
Now the KVS has increased it's scope. I am making a networked database over the simple hashmap.

### Architectural Diagram:
<img src="./dev-docs/Screenshot 2025-10-11 at 5.21.30 PM.png">

# Current Implementation
Insert:
<img src="./dev-docs/Screenshot 2025-10-20 at 3.13.57 PM.png">

Avg Latency: 2ms to 4ms [Local]

Get
<img src="./dev-docs/Screenshot 2025-10-20 at 3.14.03 PM.png">

Avg Latency: 2ms [Local]

# The Driver API
Get
```go
    kvStore.Get(key)
```

Insert
```go
    kvStore.Insert(key, value)
```

Delete
```go
    kvStore.Delete(key)
```

## My Learnings
You cannot return generic types in Go.
Either you have to have weird syntax where you take in a blank argument and return it OR you have to have a zero value and return that.


```go
// Return a rigid type [Does not make sense]
func Get[Key interfaces.Key](key Key) string {
    fmt.Println("Getting the Value for:", key)
    return "Not Yet Implemented"
}
```

```go
// Take in a blank never-used value [Does not make sense]
func Get[Key interfaces.Key, Value interfaces.Value](key Key, defaultValue Value) Value {
    fmt.Println("Getting the Value for:", key)
    return defaultValue
}
```


```go
// Return any [Does not make sense]
type Value any

func Get(key interfaces.Key) any {
    fmt.Println("Getting the Value for:", key)
    return "Not Yet Implemented"
}

```

What I did:
```go
func Get(key interfaces.Key) interface{} {
    fmt.Println("Getting the Value for:", key)
    return "Not Yet Implemented"
}

```

## Future Scope
- Use any Map as the KV store. Adding/Removing from the global KV store.
```go
    m := map[key]value
    mConverted := library.convert(&m)

    err := globalKV.Append(mConverted)
    print(err)
```

- We can just hash things twice
```go
If I have to hash: "Hello": "World";

hash(key) -> hash("Hello") = 123
hash(value) -> hash("World") = 456

globalDict {
    123: ["Hello", 456]
    456: ["World", library.END]
}
```