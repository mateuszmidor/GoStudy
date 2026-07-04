# Project Spec

## 1. Overview
- Project name: EShop
- Short description: A simplistic demo of a web shop selling products.
- Primary goal: Show golang built in features for creating REST backend
- Why this project exists: Demo project to learn Go REST API development
- Intended users: Only me
- Current stage: prototype

## 2. Success outcomes
Describe the concrete outcomes that must be true when this work is complete.

- Outcome 1: Can add a pet to the in-memory store via REST API
- Outcome 2: Can list all pets from the in-memory store via REST API
- Outcome 3: Can get details of a specific pet by ID
- Outcome 4: Can update a pet's name and price by ID

## 3. In scope
List what the agent is allowed to build/change in this iteration.

- Add Product to the Shop (POST /api/products)
- List all Products in the Shop (GET /api/products)
- Get details of given Product (GET /api/products/:id)
- Update Product details - name and price (PUT /api/products/:id)

## 4. Out of scope
Be explicit. This prevents the agent from "helpfully" expanding the feature.

- No real database - use in-memory storage only
- No authentication
- No authorization
- Minimal error handling - log error, return message in HTTP response with relevant HTTP code

## 5. User scenarios
Describe the main flows from a user/operator/developer point of view.

### Scenario 1
- As a: developer
- I want: to add a new product to the shop
- So that: the product is stored and can be retrieved later

### Scenario 2
- As a: developer
- I want: to list all products in the shop
- So that: I can see what products are available

### Scenario 3
- As a: developer
- I want: to get details of a specific product by ID
- So that: I can see the product's information

### Scenario 4
- As a: developer
- I want: to update a product's name and price
- So that: I can modify the product's information

## 6. Functional requirements
Write clear, testable requirements.

### R1 - Add Product
- Description: Add a new product to the in-memory store
- Inputs: JSON body with name (string) and price (uint, in cents)
- Expected behavior: Product is created with unique ID and stored, returns 201 Created with product data
- Edge cases: Empty name, invalid price (negative) - return 400 Bad Request
- Error handling: Log error, return 500 Internal Server Error

### R2 - List all Products
- Description: List all products from the in-memory store
- Inputs: None
- Expected behavior: Returns 200 OK with JSON array of all products
- Edge cases: Empty store - return empty array
- Error handling: Log error, return 500 Internal Server Error

### R3 - Get Product by ID
- Description: Get a specific product by its ID
- Inputs: Product ID in URL path
- Expected behavior: Returns 200 OK with product data
- Edge cases: Product not found - return 404 Not Found
- Error handling: Log error, return 500 Internal Server Error

### R4 - Update Product
- Description: Update a product's name and price
- Inputs: Product ID in URL path, JSON body with name and price
- Expected behavior: Product is updated, returns 200 OK with updated product data
- Edge cases: Product not found - return 404 Not Found; empty name or invalid price (negative) - return 400 Bad Request
- Error handling: Log error, return 500 Internal Server Error

## 7. Non-functional requirements
Define quality constraints.

- Performance: Minimal - prototype only
- Reliability: Minimal - prototype only
- Security: Not in scope
- Observability: Basic logging only
- Maintainability: Simple 2-file structure
- Portability: Cross-platform Go
- Backward compatibility: Not applicable for prototype

## 8. Technical context
- Language: Go 1.26
- Module name: eshop
- App type: REST API
- Target OS/platform: Any
- External dependencies allowed: None - use standard library only
- External dependencies forbidden: All external packages
- Storage: In-memory (map) with thread-safe access using sync.Mutex
- Network/API integrations: HTTP
- Deployment/runtime environment: Local development (port 9090)

## 9. Architecture constraints
Describe how the agent must shape the solution.

- Required architecture style: Minimal - single layer
- Package layout rules: Two files - server.go (HTTP handlers), storage.go (in-memory store + Pet model)
- Dependency direction rules: server.go depends on storage.go
- Interface boundaries: No interfaces needed for prototype
- Concurrency rules: Not required for prototype
- Configuration approach: No config - hardcoded for demo
- Error handling conventions: Log error, return HTTP error response
- Logging conventions: Standard log package to stdout
- Context propagation rules: Not required
- Cancellation/timeout rules: Not required

## 10. Existing decisions
Capture decisions already made so the agent does not re-decide them.

- We already decided to use: In-memory storage, standard library only, minimal 2-file structure
- We explicitly rejected: External database, external dependencies, complex architecture
- Naming conventions: Standard Go (camelCase for vars, PascalCase for exported types)
- API conventions: REST, JSON request/response
- Serialization format: JSON
- Testing style: Standard Go testing package
- Build/release approach: go build, go run

## 11. Repository boundaries
Tell the agent where it may and may not work.

### Allowed to modify
- storage.go
- server.go
- spec.md

### Must ask before modifying
- None

### Never modify
- None (new project)

## 12. Inputs and outputs
Define contracts clearly.

### Inputs
- Source: HTTP request body
- Format: JSON
- Validation rules: Name required (non-empty string), Price required (non-negative uint, in cents)

### Outputs
- Format: JSON
- Location: HTTP response body
- Contract/schema: Pet JSON {id: uint, name: string, price: uint}

## 13. Interfaces and contracts
Document public interfaces the agent must preserve or create.

### Public API / CLI / package contract
- Command or endpoint: REST API
- Request/input: HTTP requests to /api/pets and /api/pets/:id
- Response/output: JSON responses with appropriate HTTP status codes
- Exit codes: Not applicable
- Backward compatibility expectations: Not applicable for prototype

## 14. Data model
- Core entities: Product (id, name, price)
- Important fields: ID (uint, auto-generated starting from 1), Name (string), Price (uint, in cents)
- Invariants: ID must be unique and >= 1, Name must be non-empty, Price must be non-negative
- Validation rules: Name required, Price >= 0
- Migration needs: None

## 15. Observability
- Required logs: Error logging only
- Metrics: None
- Traces: None
- Health checks: None
- Debugging expectations: Minimal

## 16. Security and safety
- Secrets handling: Not applicable
- Authentication/authorization needs: None
- Input sanitization: Basic JSON parse validation
- Sensitive data rules: None
- Forbidden patterns: None
- Compliance constraints: None

## 17. Acceptance criteria
Use precise "WHEN / THEN" criteria.

### AC1
- WHEN: POST request is sent to /api/products with valid JSON {name: "Widget", price: 9999}
- THEN the system SHALL return 201 Created with JSON containing the product's id, name, and price (in cents) (in cents)

### AC2
- WHEN: GET request is sent to /api/products
- THEN the system SHALL return 200 OK with JSON array of all stored products

### AC3
- WHEN: GET request is sent to /api/products/:id with a valid product ID
- THEN the system SHALL return 200 OK with JSON containing the product's data

### AC4
- WHEN: GET request is sent to /api/products/:id with an invalid product ID
- THEN the system SHALL return 404 Not Found

### AC5
- WHEN: PUT request is sent to /api/products/:id with valid JSON {name: "Gadget", price: 14999}
- THEN the system SHALL return 200 OK with JSON containing the updated product's data (price in cents)

### AC6
- WHEN: PUT request is sent to /api/products/:id with an invalid product ID
- THEN the system SHALL return 404 Not Found

### AC7
- WHEN: POST request is sent with empty name or negative price
- THEN the system SHALL return 400 Bad Request

## 18. Verification plan
Tell the agent how correctness will be proven.

- Unit tests required: Yes - tests for storage and handler logic
- Integration tests required: No
- End-to-end/manual checks: Manual testing with curl/Postman
- Benchmarks/perf checks: Not required
- Lint/static analysis: go vet
- Race detection: go run -race for testing
- Coverage expectations: Decent coverage for core logic
- Done definition: All acceptance criteria pass

## 19. Execution plan
Break the work into small, reviewable chunks.

### Phase 1
- Goal: Implement PetShop REST API
- Deliverables: storage.go, server.go, basic tests

## 20. Agent working rules
Operational rules for the coding agent.

- Always explain intended changes before editing code.
- Produce a plan before implementation.
- Keep changes small and reviewable.
- Do not add dependencies without justification.
- Do not change public contracts unless explicitly allowed.
- Update tests with code changes.
- Ask before destructive refactors.
- Prefer standard library unless a dependency is justified.
- Preserve Go idioms and formatting.
- Stop and ask if the spec is ambiguous.

## 21. Commands
- Build: go build -o petshop .
- Test: go test -v ./...
- Lint: golint ./...
- Vet: go vet ./...
- Race: go run -race .
- Run locally: go run .

## 22. Deliverables
What the agent must produce.

- Code changes: storage.go, server.go
- Tests: Basic unit tests
- Documentation: This spec.md
- Migration/config updates: None
- Example usage: curl commands for testing endpoints

## 23. Open questions
Things the agent must not guess.

- None