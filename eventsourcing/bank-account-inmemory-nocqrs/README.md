# bank-account-inmemory-nocqrs

- simulates bank account; CreateAccount, FundAccount, ListAccounts, GetAccount
- uses in-memory event storage and event bus
- uses 3-tier architecture


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
| HTTP Request     | --> | Controller       | --> | Account          |
+------------------+     +------------------+     +------------------+
                                                      |
                                                      v
                                               +------------------+
                                               | Event Store      |
                                               +------------------+
```

```text
READ PATH

+------------------+     +------------------+     +------------------+
| HTTP Request     | --> | Controller       | --> | Event Store      |
+------------------+     +------------------+     +------------------+
                                                      |
                                                      v
                                               +------------------+
                                               | Rebuilt Account  |
                                               +------------------+
                                                      |
                                                      v
                                               +------------------+
                                               | HTTP Response    |
                                               +------------------+
```
