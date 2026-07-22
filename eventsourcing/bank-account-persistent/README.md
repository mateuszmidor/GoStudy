# bank-account-persistent

- simulates bank account; CreateAccount, FundAccount, ListAccounts
- uses PostgreSQL event storage and event bus for reliability

## Bug in KurrentDB eventsourcing lib
The file github.com/terraskye/eventsourcing@v0.1.6/eventstore/kurrentdb/eventstore.go has a bug:  
e.client.ReadAll reads with count=0 which means 'live reading' and newer return and the function hungs forever.  
So - using PostgreSQL event store here.

## Run

```sh
make rundb 
go run .

# 1. successfuly create and list accounts
make create
make list

# 2. successfuly create & fund & list accounts
make fund
make list # archived=false

#3. get account with balance
make get id=<uuid from the list>
```

