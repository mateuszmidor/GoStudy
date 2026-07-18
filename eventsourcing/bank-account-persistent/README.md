# bank-account-persistent

- simulates bank account; CreateAccount, FundAccount, ListAccounts
- uses KurrentDB event storage and event bus for reliability

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

