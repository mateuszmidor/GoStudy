# eventsourcing demo - task management app

Based on https://github.com/TerraSkye/eventsourcing/blob/master/docs/tutorials/index.md

## Run

```sh
go run .

# successfuly create and list tasks
make create
make list

# create task but complete it twice - error
make complete-twice

# try to complete non-existent task - error
make complete-nonexistent
```

