# sdars
Satellite radio

## 3-part application built as either:
    * 1 modular monolith
    * 3 services integrated over HTTP REST
    * 3 services integrated over GRPC
    * 3 services integrated over Kafka messages

## Could be developed by separate teams and stored in separate repositories:
    * vendor/hw
    * vendor/tuner
    * vendor/ui
    * vendor/shared_kernel
    * grpc_microservices_demo - connects hw-tuner-ui with gRPC

## Domain
Infotainment

## Subdomain
Satellite radio

## TODO
    * unit tests
    * contract tests (REST API)

## Architecture Decision Records
ADRs are created using adr tool: https://github.com/npryce/adr-tools