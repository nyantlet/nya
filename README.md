# Nya

Nya is Go version of `assert.h`.

## Example

For example, this code:

```go
x := 1 + 1
nya.Equal(3, x, "it works!")
```

will print:

```
nya: /tmp/main.go:10: main.main: it works!
expected 3 (type int), found 2 (type int)
```

## Debug assert

Nya also provides debug assertion functions.
By default, debug assertions do nothing.
You can turn on them with build tag `debugnya`.
