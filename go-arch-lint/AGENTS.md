# go-arch-lint

A small, standard-library-only Go HTTP server built as **vertical slices**.
The structure *is* the demo: each endpoint is an independent slice, and the
layout is the kind of architecture a consistency tool (e.g. go-arch-lint) would
verify. Preserve the layering when you change things.

## Layering (one-way: `cmd` → `slices` → `api`)

- `cmd/calc/main.go` — entrypoint. Builds the `*http.ServeMux` and wires each
  slice by calling its `Register`.
- `slices/{add,multiply,power}` — one slice per endpoint in its own package.
  Each exposes a public `Register(mux *http.ServeMux)` (and its own private
  handler) that mounts `"GET /<name>"`. A slice depends only on `api`, never on
  another slice.
- `api` — shared wire format + helpers only (`ParseParam`, `ReturnSuccess`,
  `ReturnError`). **All** slices use this for request parsing and JSON
  responses; don't hand-roll a wire format or import a third JSON lib.

## Commands

- `make run` (≡ `go run ./cmd/calc`) — starts the server on **:9090**.
- There is **no** test/lint/vet target and no test suite. Verify with
  `go build ./...`, `go vet ./...`, `gofmt -l .`.
- Endpoints: `GET /add`, `/multiply`, `/power` — all require `a` and `b`
  (float64). `200 {"result":<n>}` on success; `400 {"error":"..."}` on bad/missing
  args (`/power` also 400s on NaN/Inf).

## Adding a new slice

1. `slices/<name>/<name>.go`, package `<name>`, exposing
   `func Register(mux *http.ServeMux)` that mounts `"GET /<name>"`.
2. Parse with `api.ParseParam`, reply with `api.ReturnSuccess` / `api.ReturnError`.
3. Import it in `cmd/calc/main.go` and call `Register(mux)`.

## Constraints

- Go 1.26+. **Standard library only** — no external modules (no `go.sum`); keep
  it that way.
- Respect the one-way dependency (`cmd` → `slices` → `api`). Introducing
  cross-slice imports or a second shared package breaks the point of the demo.
