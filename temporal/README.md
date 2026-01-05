# temporal.io

Temporal allows to model a long running business process into a resilient, reliable and durable workflow, e.g. a Mortgage.

## Components of a Temporal solution

- Activities - business process steps that can take long time (weeks) or unexpectedly fail (external service call failure, application crash)
- Workflow - orchiestrates steps into a business process, where each step can be retried in case of a failure
- Application - requests an execution of a workflow with given params
  - sends the request to Temporal service, which places it on a queue
- Worker - handles execution of a process
  - registers in Temporal service and polls selected queue for requests to handle

## install CLI

Linux: https://temporal.download/cli/archive/latest?platform=linux&arch=amd64

, or
```
brew install temporal
```

## run cluster

```
git clone https://github.com/temporalio/docker-compose
cd docker-compose
docker-compose up
```

Dashboard: http://localhost:8080