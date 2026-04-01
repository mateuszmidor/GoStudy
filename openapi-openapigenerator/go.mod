module github.com/mateuszmidor/GoStudy/openapi-openapigenerator

go 1.18

require github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_server v0.0.0-00010101000000-000000000000

require github.com/gorilla/mux v1.8.0 // indirect

replace github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_server => ./generated_server

replace github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_client => ./generated_client
