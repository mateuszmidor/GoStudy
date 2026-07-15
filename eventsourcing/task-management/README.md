# eventsourcing demo - task management app

Based on https://github.com/TerraSkye/eventsourcing/blob/master/docs/tutorials/index.md

## Run

```sh
go run .

# 1. successfuly create and list tasks
make create
make list

# 2. successfuly create & complete & after 5sec archive task
make complete
make list # archived=false
# wait 5sec
make list # archived=true

# 3. create task but complete it twice - error
make complete-twice

# 4. try to complete non-existent task - error
make complete-nonexistent
```

