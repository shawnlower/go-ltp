# go-ltp
golang CLI for LTP

# Goals by version:

0.0.1: Learn basic golang syntax; make something that doesn't kill kittens
    - Parse command-line
    - Read data from stdin
    - Generate hash from stream
    - Create object with:
        - timestamp
        - hash
        - contents
        - stdin socket's peer (if available in kernel; req 3.3+)

0.0.2: Do something with our input
    - Keep database. e.g.: ./data/database.sqlite
    - Write to file, e.g.: ./data/11c52179958394840

0.0.3: File support
    - ltp ./myfile.png
    - 

0.0.4: Use the webs
    - Post meta-data
    - Post object to HTTP endpoint

# Notes by version:

## 0.0.1: Write a basic CLI in Go

### CLI framework: cobra
The [cobra (github)] package is used by kubectl, docker, openshift, and etcd


### Config parser: viper

### Interface

ltp
    add [pathspec|uri|-]: add an item
    list:                 list all objects
    get x:                get data for object 'x'
    show x:               get meta-data for object 'x'
    del x:                delete object 'x'
    edit x:               edit object 'x' (links, etc)

```
# Fetch and build cobra CLI package
go get github.com/spf13/cobra/cobra
```

```
$ cobra init github.com/shawnlower/go-ltp/go-ltp
Your Cobra application is ready at
/home/shawn/projects/go-ltp/src/github.com/shawnlower/go-ltp/go-ltp
```

## Digest and io.Reader jobs based on:
https://github.com/docker/distribution/blob/749f6afb4572201e3c37325d0ffedb6f32be8950/cmd/digest/main.go

viper for configs



# References
[cobra (github)]: https://github.com/spf13/cobra 
