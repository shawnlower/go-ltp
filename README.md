# go-ltp
golang CLI for LTP

Goals by version:

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
