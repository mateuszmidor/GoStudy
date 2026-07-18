# bank-account-persistent

- simulates bank account; CreateAccount, FundAccount, ListAccounts
- uses KurrentDB event storage and event bus for reliability

## Bug in eventsourcing lib
The file github.com/terraskye/eventsourcing@v0.1.6/eventstore/kurrentdb/eventstore.go has a bug:  
e.client.ReadAll reads with count=0 which means newer return and the function hungs forever.

## Run

```sh
make rundb # KurrentDB Admin UI: http://localhost:2113
go run .

# 1. successfuly create and list accounts
make create
make list

# 2. successfuly create & fund & list accounts
make fund
make list # archived=false
```

