# only create, 1 account
create:
	curl -X POST http://localhost:8080/accounts \
	-H "Content-Type: application/json" \
	-d '{"owner_name": "Andrzej"}'

fund:
	ACCOUNT_ID=$$(curl -X POST "http://localhost:8080/accounts" -H "Content-Type: application/json" -d '{"owner_name": "Andrzej"}' | jq '.account_id') && \
	curl -X POST "http://localhost:8080/accounts/$$ACCOUNT_ID/atm"  -H "Content-Type: application/json" -d '{"Dollars": 5}' && \
	curl "http://localhost:8080/accounts"

# just list accounts
list:
	@curl http://localhost:8080/accounts

