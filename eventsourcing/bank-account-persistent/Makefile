# only create, 1 account
create:
	@curl -X POST http://localhost:8080/accounts \
	-H "Content-Type: application/json" \
	-d '{"owner_name": "Andrzej"}'

# create and fund account
fund:
	@ACCOUNT_ID=$$(curl -s -X POST "http://localhost:8080/accounts" -H "Content-Type: application/json" -d '{"owner_name": "Andrzej"}' | jq '.account_id') && \
	curl -s -X POST "http://localhost:8080/accounts/$$ACCOUNT_ID/atm"  -H "Content-Type: application/json" -d '{"Dollars": 5}' && \
	curl -s "http://localhost:8080/accounts"

# just list accounts
list:
	@curl http://localhost:8080/accounts

rundb:
	docker run --rm -d -p 2113:2113 -e EVENTSTORE_INSECURE=true kurrentplatform/kurrentdb:latest --insecure