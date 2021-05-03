
# CRUD App Example

> A CRUD application which edits todo records in Command Golang. 

## Build Setup

```bash

# build
$ make build

# get list todo
$ make list
$ ./main list

# create todo
$ ./main create "Todo text"

# update todo
$ ./main update {uuid} "Todo text"

# delete todo
$ ./main create {uuid}

# mark done todo
$ ./main mark-done {uuid}
```
