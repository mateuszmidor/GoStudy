# create account
create:
	@curl -X POST http://localhost:8080/accounts \
	-H "Content-Type: application/json" \
	-d '{"owner_name": "Andrzej"}'

# create and fund account with 5 dollars
fund:
	@ACCOUNT_ID=$$(curl -s -X POST "http://localhost:8080/accounts" -H "Content-Type: application/json" -d '{"owner_name": "Dorian"}' | jq '.account_id') && \
	curl -s -X POST "http://localhost:8080/accounts/$$ACCOUNT_ID/deposits"  -H "Content-Type: application/json" -d '{"Dollars": 5}' && \
	curl -s "http://localhost:8080/accounts"

# list accounts
list:
	@curl http://localhost:8080/accounts

# get account balance, example: make get id=bf5b9321-eb0f-4089-888e-45c8a9d5acd2
get:
	@curl http://localhost:8080/accounts/$(id)/balance

# run the event store db
rundb:
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:16-alpine