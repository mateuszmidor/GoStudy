up:
	docker run --name jaeger \
	-d \
	--rm \
	-e COLLECTOR_OTLP_ENABLED=true \
	-p 16686:16686 \
	-p 4317:4317 \
	-p 4318:4318 \
	jaegertracing/all-in-one:1.35

run:
	go run .

down:
	docker stop jaeger

houses:
	while true; do curl -s http://localhost:8080/build-house; sleep 3; done