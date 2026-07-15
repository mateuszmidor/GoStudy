# only create, 2 tasks
create:
	curl -X POST http://localhost:8080/api/v1/tasks \
	-H "Content-Type: application/json" \
	-d '{"title": "Task One"}'

	curl -X POST http://localhost:8080/api/v1/tasks \
	-H "Content-Type: application/json" \
	-d '{"title": "Task Two"}'

# create and complete, entails archiving it after 5s delay
complete:
	TASK_ID=$$(curl -s -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"title": "Finish tutorial"}' | jq -r '.task_id') && \
	curl -X POST "http://localhost:8080/api/v1/tasks/$$TASK_ID/complete"

# just list tasks
list:
	@curl http://localhost:8080/api/v1/tasks

# cause error due to business rule violation - can't complete single task twice
complete-twice:
	TASK_ID=$$(curl -s -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"title": "Finish tutorial"}' | jq -r '.task_id') && \
	curl -X POST "http://localhost:8080/api/v1/tasks/$$TASK_ID/complete" && \
	curl -X POST "http://localhost:8080/api/v1/tasks/$$TASK_ID/complete"

# cause error due to business rule violation - can't complete non existent task
complete-nonexistent:
	curl -X POST "http://localhost:8080/api/v1/tasks/00000000-0000-0000-0000-000000000000/complete"