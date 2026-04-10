default:
	# configure the connector "process_list" to ingest file "/tmp/process_list.txt" mounted as docker volume
	curl -X PUT http://localhost:8083/connectors/process_list/config \
		-H 'Content-Type: application/json' \
		--data @connect-request.json

	# populate the source file
	ps aux >> process_list.txt