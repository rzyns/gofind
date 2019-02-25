Find a symbol and output `file:#offset` position for use in `guru`.

## Example ##

```

# Raw output
$ gofind ./main main
/Users/user1/git/gofind/main.go:#81

# Used with guru
$ guru -reflect what "$(go run main.go . main)"
/Users/user1/git/gofind/main.go:10.6-10.9: identifier
/Users/user1/git/gofind/main.go:10.1-52.1: function declaration
/Users/user1/git/gofind/main.go:1.1-52.1: source file
-: modes: [callers callstack definition describe implements pointsto referrers whicherrs]
-: srcdir:
-: import path:
/Users/user1/git/gofind/main.go:10:6: main
```
