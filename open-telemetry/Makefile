jaeger-up:
	docker run --name jaeger \
	-d \
	-e COLLECTOR_OTLP_ENABLED=true \
	-p 16686:16686 \
	-p 4317:4317 \
	-p 4318:4318 \
	jaegertracing/all-in-one:1.35

run:
	go run .

jaeger-down:
	docker stop jaeger