# go-ltp
golang CLI for LTP

# Goals by version:

## TODO

### Oct 1-31

0.0.4: Enable creation of new items/terms on remote server (Use the webs)
    - NewItem(), ShowItem() (or similar) methods

0.0.5:
    - GCS or S3 store for objects

0.0.6:
    - Link items / objects

0.0.7:
    - Remote parsing/indexing triggered via pubsub/etc

### Future

- test testing
- Enable HTTP API
- Connect ltpweb-basic to fetch data
- Auth (!)

## DONE.

### Sept 1 - Sept 30
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

# Notes by version:

## 0.0.1: Write a basic CLI in Go

### CLI framework: cobra
The [cobra (github)] package is used by kubectl, docker, openshift, and etcd


### Config parser: viper

### Interface

ltp
    add [pathspec|uri|-]: add an item
    list:                 list all objects
    get x:                get data for object 'x' to a file
    cat x:                get, but to stdout; can use alt-formats
                          e.g. cat --format=html
    show x:               get meta-data for object 'x'
    del|rm x:             delete object 'x'
    edit x:               edit object 'x' (links, etc)
    link|ln               link items

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

# Notes [methods]

## ADD method
ltp add [FILE | -]

0. The client reads the source data, extracts meta-data, generates a per-object key
0. The object is uploaded to an object store
0. The URL for the object, as well as the per-object key is stored in the database
0. The metadata for the object is uploaded

## SHOW method
ltp show [term | id]

## GET method

# Notes [namespaces]

## Creating a thing

Better yet, let's make a place
```
$ ltp new -t schema:Place
Created: http://shawnlower.net/i/home

$ ltp show home
Item: http://shawnlower.net/i/home
Type: http://schema.org/Thing
```

See prev notes on observations.

Use UUIDv1 + namespace as the quad's label; e.g.
7309d2c8-5082-40d4-9d73-a81ad1ed9d14.v1.client.ltp.shawnlower.net

```
uuid, err := uuid.NewUUID()
if err != nil {
	log.Error("Failed to get UUID: ", err)
}

log.Debug("Got a UUID: ", uuid)
log.Debug("uuid.ClockSequence: ", uuid.ClockSequence())
log.Debug("uuid.Domain: ", uuid.Domain())
log.Debug("uuid.ID: ", uuid.ID())
log.Debug("uuid.NodeID: ", uuid.NodeID())
log.Debug("uuid.String: ", uuid.String())
log.Debug("uuid.Time: ", uuid.Time())
log.Debug("uuid.URN: ", uuid.URN())
log.Debug("uuid.Variant: ", uuid.Variant())
log.Debug("uuid.Version: ", uuid.Version())
```

Required info:
- Name
- Creator/Client ID (used as label in quad/assertion)
Optional:
- Type (default <schema:Thing>)
Inferred:

Becomes:
<ltp:home> <a> <schema:Place> <uuid.v1.client.ltp.shawnlower.net>
<uuid.v1.client.ltp.shawnlower.net> [ provenance statement(s) ]

# List of places:
See first: rdf container (open/size undefined) vs rdf collection (closed/fixed-size)
Also: schema.org/ItemList

# References
[cobra (github)]: https://github.com/spf13/cobra 
