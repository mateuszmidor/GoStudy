# testcontainers

Run Postgresql in docker container programatically to test the Repository implementation.

## Run

First run the Docker daemon, then:  
```sh
go test -count=1 -v ./...
```

```log
=== RUN   TestCustomerRepoTestSuite
2025/12/18 13:39:57 running postgresql container
2025/12/18 13:39:57 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 24.0.2
  API Version: 1.43
  Operating System: Docker Desktop
  Total Memory: 3932 MB
  Testcontainers for Go Version: v0.40.0
  Resolved Docker Host: unix:///var/run/docker.sock
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: aabdd008e854c8b7c8a18e5119cdddd6fd036a55fc5c81a398c0747d0cf9a4dd
  Test ProcessID: 444c6303-2f5a-4466-8764-a36cf3d05667
2025/12/18 13:39:57 🐳 Creating container for image postgres:15.3-alpine
2025/12/18 13:39:57 🐳 Creating container for image testcontainers/ryuk:0.13.0
2025/12/18 13:39:57 ✅ Container created: 4fc891835fe3
2025/12/18 13:39:57 🐳 Starting container: 4fc891835fe3
2025/12/18 13:39:57 ✅ Container started: 4fc891835fe3
2025/12/18 13:39:57 ⏳ Waiting for container id 4fc891835fe3 image: testcontainers/ryuk:0.13.0. Waiting for: port 8080/tcp to be listening
2025/12/18 13:39:58 Shell not executable in container, only external port validated
2025/12/18 13:39:58 🔔 Container is ready: 4fc891835fe3
2025/12/18 13:39:58 ✅ Container created: ce5e7ea47800
2025/12/18 13:39:58 🐳 Starting container: ce5e7ea47800
2025/12/18 13:39:58 ✅ Container started: ce5e7ea47800
2025/12/18 13:39:58 ⏳ Waiting for container id ce5e7ea47800 image: postgres:15.3-alpine. Waiting for: all of: [log message "database system is ready to accept connections" (occurrence: 2)]
2025/12/18 13:40:01 🔔 Container is ready: ce5e7ea47800
=== RUN   TestCustomerRepoTestSuite/TestCreateCustomer
=== RUN   TestCustomerRepoTestSuite/TestGetCustomerByEmail
2025/12/18 13:40:01 stopping postgresql container
2025/12/18 13:40:01 🐳 Stopping container: ce5e7ea47800
2025/12/18 13:40:01 ✅ Container stopped: ce5e7ea47800
2025/12/18 13:40:01 🐳 Terminating container: ce5e7ea47800
2025/12/18 13:40:02 🚫 Container terminated: ce5e7ea47800
--- PASS: TestCustomerRepoTestSuite (5.06s)
    --- PASS: TestCustomerRepoTestSuite/TestCreateCustomer (0.01s)
    --- PASS: TestCustomerRepoTestSuite/TestGetCustomerByEmail (0.01s)
PASS
ok      github.com/mateuszmidor/GoStudy/db/testcontainers       5.645s
```