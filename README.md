# Go Practice — Revision Index

One Go module. Every program lives in its own folder as `main.go`. Run any of them from the repo root:

```bash
go run ./05-dsa/05-blind-75/01-two-sum
go build ./...   # check everything compiles
```

## 01-basics — language fundamentals
| # | Topic | # | Topic |
|---|---|---|---|
| 01 | hello | 09 | slices (`append`, `sort`) |
| 02 | variables, types, `:=` | 10 | maps |
| 03 | user input (`bufio`) | 11 | structs, methods |
| 04 | conversion (`strconv`) | 12 | interfaces |
| 05 | time | 13 | enums (`iota`) |
| 06 | pointers | 14 | generics |
| 07 | arrays | 15 | files (`os`, `io`) |
| 08 | functions (multi-return, variadic, anonymous) | | |

## 02-concurrency
- `01-goroutines`: goroutines + WaitGroup + Mutex, checking site status
- `02-mutex-waitgroup`: race condition fix
- `03-channels`: buffered → select → for-select → worker pool → done channel → pipelines → timeouts → non-blocking → range → reading a closed channel

## 03-web
http GET → URL parsing → http client → JSON → gorilla/mux → REST API → auth middleware → MongoDB API

## 04-devops-tasks — interview-style practical tasks
| Folder | Task |
|---|---|
| 01-yaml-update | Edit a multi-doc YAML (`yaml.Node`) |
| 02-crd-yaml | Sample CRD + instance |
| 03-gitlab-client | GitLab API client (⚠ **deletes** a project if run) |
| 04-url-query-param | Pull a query param out of a URL |
| 05-clusters-endpoint | GET `/clusters` with filters |
| 06-url-builder | Build a URL with query params |
| 07-k8s-webhook | Webhook receiver |
| 08-word-frequency | Count words |
| 09-log-parser | Parse logs over HTTP |
| 10-api-get | HTTP GET + JSON with context timeout |
| 11-api-md5 | Concurrent URL fetch + md5 |

## 05-dsa
- `01-patterns`: star/number/alphabet patterns (nested loops)
- `02-math`: count digits, reverse, palindrome, armstrong, divisors, prime, power of two, factorial
- `03-recursion`: basics + notes
- `04-arrays`: min/max, unique, intersection, XOR single number, longest consecutive, sliding window
- `05-blind-75`: two-sum, stock, product-except-self, duplicates, top-k, anagrams, parentheses, longest substring, clone graph, islands, tree depth, palindrome

## notes/
| File | What |
|---|---|
| `go-notes.md` | Structs, memory, go mod, goroutines, channels |
| `revision-notes.md` | One-liners for quick revision |
| `dsa-cheat-sheet.md` | Daily coding fundamentals cheat sheet |
| `big-o.md` | Complexity cheat sheet |
| `crd.md` | CRD, RBAC, controllers, operators |
