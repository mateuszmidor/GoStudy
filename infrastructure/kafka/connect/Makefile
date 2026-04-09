default:
	ps aux  > /tmp/kernel.log
	curl -X PUT http://localhost:8083/connectors/process/config \
		-H 'Content-Type: application/json' \
		--data @connect-request.json