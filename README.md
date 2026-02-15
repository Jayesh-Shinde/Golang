# Go Learning — Personal Exercises & Examples

A curated collection of small Go exercises, practice projects and examples used for learning. Each folder contains self-contained examples (many with `go.mod`) you can run locally to practice core Go concepts.

---

## Quick overview 🔎

- **Language / version:** most modules use `go 1.25.5` (see individual `go.mod`).
- **How to run:** either `go run <file.go>` for single-file examples or `cd <folder> && go run .` for module folders.

---

## Folders & what they contain 📂

- `arrays/` — examples for arrays, slices and capacity/copy behavior.
  - Run: `go run arrays/lists.go`

- `basic-examples/` — small standalone examples (e.g. `investment_calculator.go`).
  - Run: `go run basic-examples/investment_calculator.go`

- `basic-practice/` — short practice exercises (profit calculator).
  - Run: `go run basic-practice/profit_calculator.go`

- `concurrency/` — goroutines & channels demos.
  - Run: `go run concurrency/main.go`

- `control-structure/` — branching/loops + simple interactive "bank" app (uses internal `fileops`).
  - Run: `cd control-structure && go run .`

- `DSA/` — data-structures & algorithm problems (e.g. Best Time to Buy Stock).
  - Run: `go run DSA/bestTimeToBuyStock/main.go`

- `functions/` — examples for function types, variadic and anonymous functions.
  - Run: `go run functions/variadic/main.go` or `go run functions/anonymous/anonymous.go`
  - Note: `functions/function/` contains reusable function code (not necessarily a runnable `package main`).

- `maps/` — map usage, iteration and examples.
  - Run: `go run maps/map.go`

- `pointers/` — pointer basics & pointer receivers.
  - Run: `go run pointers/pointer.go`

- `practiceProject/` — small multi-file project that reads `prices.txt` and writes JSON results; demonstrates concurrency, file I/O and modular code.
  - Run: `cd practiceProject && go run .`
  - Output: `result_*.json` files are produced by jobs.

- `Rest-Api/` — small REST API (Gin + SQLite) demonstrating routing, middleware, JWT and persistence.
  - Run: `cd Rest-Api && go run .` → server at `http://localhost:8080`
  - Endpoints: `GET /events`, `POST /signup`, `POST /login`, protected routes for creating/updating events.
  - DB: `api.db` (SQLite) will be created automatically.

- `struct-practice/` — working with structs, interfaces and basic I/O (creates `todo.json` / note files).
  - Run: `cd struct-practice && go run .`

- `structs-custom-types/01-starting-project/` — example of custom types and constructors (uses Go module `go 1.21.2`).
  - Run: `cd structs-custom-types/01-starting-project && go run .`

---

## Recommended setup 🔧

1. Install Go (recommended: same major version as `go.mod`; this workspace mostly uses **Go 1.25.5**).
2. From repository root you can run individual examples or `cd` into module folders and run `go run .`.

Examples:

```bash
# run a one-file example
go run arrays/lists.go

# run a module
cd Rest-Api && go run .
```

---

## Tips & notes 💡

- Check each folder's `go.mod` if you need the exact Go version or module path.
- `Rest-Api` uses SQLite (`api.db` will appear after running).
- `practiceProject` reads `prices.txt` and writes JSON outputs (see `result_*.json`).

---

## Want to push this to a Git repo? 🚀

```bash
cd /path/to/Go
git init
git add .
git commit -m "Initial: Go learning workspace"
# add remote (replace URL) and push
git remote add origin git@github.com:youruser/go-learning.git
git branch -M main
git push -u origin main
```

---

## Contributing / next steps ✅

- Add new examples in new folders (follow existing naming and add `go.mod` if it's a standalone module).
- Add unit tests where appropriate (`*_test.go`).
- Add a `LICENSE` if you plan to publish.

---

If you want, I can:

- add a `LICENSE` file, or
- create GitHub Actions to run `go test` on pushes, or
- convert the workspace into a single monorepo with a top-level `go.work`.

Pick any one and I will add it for you. 👇
