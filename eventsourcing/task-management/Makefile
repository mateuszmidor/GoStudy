create:
	curl -X POST http://localhost:8080/api/v1/tasks \
	-H "Content-Type: application/json" \
	-d '{"title": "Task One"}'

	curl -X POST http://localhost:8080/api/v1/tasks \
	-H "Content-Type: application/json" \
	-d '{"title": "Task Two"}'

list:
	@curl http://localhost:8080/api/v1/tasks

complete-twice:
	TASK_ID=$$(curl -s -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"title": "Finish tutorial"}' | jq -r '.task_id') && \
	curl -X POST "http://localhost:8080/api/v1/tasks/$$TASK_ID/complete" && \
	curl -X POST "http://localhost:8080/api/v1/tasks/$$TASK_ID/complete"

complete-nonexistent:
	curl -X POST "http://localhost:8080/api/v1/tasks/00000000-0000-0000-0000-000000000000/complete"