# bank-account-inmemory

- simulates bank account; CreateAccount, FundAccount, ListAccounts, GetAccount
- uses in-memory event storage and event bus
- uses CQRS architectural style

## Run

```sh
go run .

# 1. successfuly create and list accounts
make create
make list 

# 2. successfuly create & fund & list accounts
make fund
make list # archived=false

# 3. get account
make get id=<uuid from the list>
```

## How it works

```text
WRITE PATH

+------------------+     +------------------+     +------------------+
| HTTP Request     | --> | HTTP Handler     | --> | Command Use Case |
+------------------+     +------------------+     +------------------+
                                                       |
                                                       v
                                                +------------------+
                                                | Event Store      |
                                                +------------------+
                                                       |
                                                       v
                                                +------------------+
                                                | Account State    |
                                                +------------------+
```

```text
READ PATH

+------------------+     +------------------+     +------------------+
| HTTP Request     | --> | HTTP Handler     | --> | Query Use Case   |
+------------------+     +------------------+     +------------------+
                                                       |
                                                       v
                                                +------------------+
                                                | Event Store      |
                                                +------------------+
                                                       |
                                                       v
                                                +------------------+
                                                | Rebuilt View     |
                                                +------------------+
                                                       |
                                                       v
                                                +------------------+
                                                | HTTP Response    |
                                                +------------------+
```
