default:
	sudo dmesg > /tmp/kernel.log
	curl -X PUT http://localhost:8083/connectors/kernel-logs/config \
		-H 'Content-Type: application/json' \
		--data @connect-request.json