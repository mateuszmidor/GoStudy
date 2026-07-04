# Project Spec

## 1. Overview
- Project name:
- Short description:
- Primary goal:
- Why this project exists:
- Intended users:
- Current stage: new project | existing codebase | prototype | production hardening

## 2. Success outcomes
Describe the concrete outcomes that must be true when this work is complete.

- Outcome 1:
- Outcome 2:
- Outcome 3:

## 3. In scope
List what the agent is allowed to build/change in this iteration.

-
-
-

## 4. Out of scope
Be explicit. This prevents the agent from “helpfully” expanding the feature.

-
-
-

## 5. User scenarios
Describe the main flows from a user/operator/developer point of view.

### Scenario 1
- As a:
- I want:
- So that:

### Scenario 2
- As a:
- I want:
- So that:

## 6. Functional requirements
Write clear, testable requirements.

### R1
- Description:
- Inputs:
- Expected behavior:
- Edge cases:
- Error handling:

### R2
- Description:
- Inputs:
- Expected behavior:
- Edge cases:
- Error handling:

## 7. Non-functional requirements
Define quality constraints.

- Performance:
- Reliability:
- Security:
- Observability:
- Maintainability:
- Portability:
- Backward compatibility:

## 8. Technical context
- Language: Go version
- Module name:
- App type: CLI | API | worker | TUI | library | service
- Target OS/platform:
- External dependencies allowed:
- External dependencies forbidden:
- Storage:
- Network/API integrations:
- Deployment/runtime environment:

## 9. Architecture constraints
Describe how the agent must shape the solution.

- Required architecture style: e.g. hexagonal, layered, clean architecture
- Package layout rules:
- Dependency direction rules:
- Interface boundaries:
- Concurrency rules:
- Configuration approach:
- Error handling conventions:
- Logging conventions:
- Context propagation rules:
- Cancellation/timeout rules:

## 10. Existing decisions
Capture decisions already made so the agent does not re-decide them.

- We already decided to use:
- We explicitly rejected:
- Naming conventions:
- API conventions:
- Serialization format:
- Testing style:
- Build/release approach:

## 11. Repository boundaries
Tell the agent where it may and may not work.

### Allowed to modify
-
-
-

### Must ask before modifying
-
-
-

### Never modify
-
-
-

## 12. Inputs and outputs
Define contracts clearly.

### Inputs
- Source:
- Format:
- Validation rules:

### Outputs
- Format:
- Location:
- Contract/schema:

## 13. Interfaces and contracts
Document public interfaces the agent must preserve or create.

### Public API / CLI / package contract
- Command or endpoint:
- Request/input:
- Response/output:
- Exit codes:
- Backward compatibility expectations:

## 14. Data model
- Core entities:
- Important fields:
- Invariants:
- Validation rules:
- Migration needs:

## 15. Observability
- Required logs:
- Metrics:
- Traces:
- Health checks:
- Debugging expectations:

## 16. Security and safety
- Secrets handling:
- Authentication/authorization needs:
- Input sanitization:
- Sensitive data rules:
- Forbidden patterns:
- Compliance constraints:

## 17. Acceptance criteria
Use precise “WHEN / THEN” criteria.

### AC1
- WHEN:
- THEN the system SHALL:

### AC2
- WHEN:
- THEN the system SHALL:

### AC3
- WHEN:
- THEN the system SHALL:

## 18. Verification plan
Tell the agent how correctness will be proven.

- Unit tests required:
- Integration tests required:
- End-to-end/manual checks:
- Benchmarks/perf checks:
- Lint/static analysis:
- Race detection:
- Coverage expectations:
- Done definition:

## 19. Execution plan
Break the work into small, reviewable chunks.

### Phase 1
- Goal:
- Deliverables:

### Phase 2
- Goal:
- Deliverables:

### Phase 3
- Goal:
- Deliverables:

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
- Build:
- Test:
- Lint:
- Vet:
- Race:
- Run locally:

## 22. Deliverables
What the agent must produce.

- Code changes:
- Tests:
- Documentation:
- Migration/config updates:
- Example usage:

## 23. Open questions
Things the agent must not guess.

-
-
-